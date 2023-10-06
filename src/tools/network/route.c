#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netinet/ip.h>
#include <netinet/ip_icmp.h>
#include <netdb.h>
#include <signal.h>
#include <sys/time.h>

#define MAX_HOPS 30
#define PACKET_SIZE 60
#define TIMEOUT 1

unsigned short in_cksum(unsigned short *addr, int len) {
    int nleft = len;
    int sum = 0;
    unsigned short *w = addr;
    unsigned short answer = 0;

    while (nleft > 1) {
        sum += *w++;
        nleft -= 2;
    }

    if (nleft == 1) {
        *(unsigned char *)(&answer) = *(unsigned char *)w;
        sum += answer;
    }

    sum = (sum >> 16) + (sum & 0xffff);
    sum += (sum >> 16);
    answer = ~sum;
    return answer;
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <host>\n", argv[0]);
        return 1;
    }

    char *host = argv[1];
    printf("traceroute to %s (%s), %d hops max, %d byte packets\n",
           host, host, MAX_HOPS, PACKET_SIZE);

    struct sockaddr_in target;
    memset(&target, 0, sizeof(struct sockaddr_in));
    target.sin_family = AF_INET;

    struct hostent *host_info = gethostbyname(host);
    if (!host_info) {
        perror("gethostbyname");
        return 1;
    }
    memcpy(&target.sin_addr, host_info->h_addr, host_info->h_length);

    for (int ttl = 1; ttl <= MAX_HOPS; ttl++) {
        int sockfd = socket(AF_INET, SOCK_RAW, IPPROTO_ICMP);
        if (sockfd == -1) {
            perror("socket");
            return 1;
        }

        if (setsockopt(sockfd, SOL_IP, IP_TTL, &ttl, sizeof(ttl)) == -1) {
            perror("setsockopt");
            return 1;
        }

        struct timeval timeout;
        timeout.tv_sec = TIMEOUT;
        timeout.tv_usec = 0;
        if (setsockopt(sockfd, SOL_SOCKET, SO_RCVTIMEO, &timeout, sizeof(timeout)) == -1) {
            perror("setsockopt");
            return 1;
        }

        struct icmp icmp_packet;
        memset(&icmp_packet, 0, sizeof(struct icmp));
        icmp_packet.icmp_type = ICMP_ECHO;
        icmp_packet.icmp_code = 0;
        icmp_packet.icmp_id = getpid();
        icmp_packet.icmp_seq = ttl;
        icmp_packet.icmp_cksum = 0;
        icmp_packet.icmp_cksum = in_cksum((unsigned short *)&icmp_packet, sizeof(struct icmp));

        ssize_t bytes_sent = sendto(sockfd, &icmp_packet, sizeof(struct icmp), 0,
                                    (struct sockaddr *)&target, sizeof(struct sockaddr_in));
        if (bytes_sent == -1) {
            perror("sendto");
            return 1;
        }

        struct sockaddr_in reply_addr;
        socklen_t reply_len = sizeof(struct sockaddr_in);
        char reply_buffer[PACKET_SIZE];

        if (recvfrom(sockfd, reply_buffer, sizeof(reply_buffer), 0,
                     (struct sockaddr *)&reply_addr, &reply_len) == -1) {
            printf("%2d  *\n", ttl);
        } else {
            struct ip *ip_reply = (struct ip *)reply_buffer;
            struct icmp *icmp_reply = (struct icmp *)(reply_buffer + (ip_reply->ip_hl << 2));

            printf("%2d  %s  %.3f ms\n", ttl,
                   inet_ntoa(reply_addr.sin_addr),
                   (double)(icmp_reply->icmp_id == getpid() ? 1 : 0) * TIMEOUT * 1000);
        }

        close(sockfd);
    }

    return 0;
}