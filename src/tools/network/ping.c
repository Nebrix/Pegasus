#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <netinet/ip_icmp.h>
#include <netinet/in.h>
#include <sys/time.h>

#define PACKET_SIZE 64
#define MAX_PACKET_SIZE 1024
#define MAX_TRIES 4

unsigned short checksum(void *b, int len) {
    unsigned short *buf = b;
    unsigned int sum = 0;
    unsigned short result;

    for (sum = 0; len > 1; len -= 2)
        sum += *buf++;

    if (len == 1)
        sum += *(unsigned char *)buf;

    sum = (sum >> 16) + (sum & 0xFFFF);
    sum += (sum >> 16);
    result = ~sum;

    return result;
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        printf("Usage: %s <hostname or IP address>\n", argv[0]);
        return 1;
    }

    char *hostname = argv[1];
    struct sockaddr_in addr;
    struct icmphdr icmp_header;
    char packet[MAX_PACKET_SIZE];

    int sockfd = socket(AF_INET, SOCK_RAW, IPPROTO_ICMP);
    if (sockfd < 0) {
        perror("socket");
        return 1;
    }

    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    inet_pton(AF_INET, hostname, &addr.sin_addr);

    int seq = 0;
    for (int tries = 0; tries < MAX_TRIES; tries++) {
        memset(&icmp_header, 0, sizeof(icmp_header));
        icmp_header.type = ICMP_ECHO;
        icmp_header.code = 0;
        icmp_header.un.echo.id = getpid();
        icmp_header.un.echo.sequence = seq++;
        icmp_header.checksum = 0;
        icmp_header.checksum = checksum(&icmp_header, sizeof(icmp_header));

        memcpy(packet, &icmp_header, sizeof(icmp_header));

        struct timeval start_time, end_time;
        gettimeofday(&start_time, NULL);

        if (sendto(sockfd, packet, sizeof(icmp_header), 0, (struct sockaddr *)&addr, sizeof(addr)) <= 0) {
            perror("sendto");
            return 1;
        }

        socklen_t addr_len = sizeof(addr);
        unsigned char buffer[PACKET_SIZE];

        if (recvfrom(sockfd, buffer, sizeof(buffer), 0, (struct sockaddr *)&addr, &addr_len) <= 0) {
            perror("recvfrom");
            return 1;
        }

        gettimeofday(&end_time, NULL);

        double elapsed_time = (end_time.tv_sec - start_time.tv_sec) * 1000.0 + (end_time.tv_usec - start_time.tv_usec) / 1000.0;
        printf("Received ICMP reply from %s: seq=%d time=%.2f ms\n", hostname, seq - 1, elapsed_time);

        sleep(1);
    }

    close(sockfd);
    return 0;
}