#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <netdb.h>

#define WHOIS_SERVER "whois.iana.org"
#define WHOIS_PORT "43"

void print_error(const char* msg) {
    fprintf(stderr, "Error: %s\n", msg);
    exit(1);
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <domain>\n", argv[0]);
        exit(1);
    }

    const char *domain = argv[1];

    int sockfd;
    struct addrinfo hints, *res;

    memset(&hints, 0, sizeof(hints));
    hints.ai_family = AF_UNSPEC;
    hints.ai_socktype = SOCK_STREAM;

    if (getaddrinfo(WHOIS_SERVER, WHOIS_PORT, &hints, &res) != 0) {
        print_error("socket");
    }
    
    sockfd = socket(res->ai_family, res->ai_socktype, res->ai_protocol);
    if (sockfd == -1) {
        print_error("socket");
    }

    if (connect(sockfd, res->ai_addr, res->ai_addrlen) == -1) {
        print_error("connect");
    }

    freeaddrinfo(res);

    char request[256];
    snprintf(request, sizeof(request), "%s\r\n", domain);
    if (send(sockfd, request, strlen(request), 0) == -1) {
        print_error("send");
    }

    char response[4096];
    ssize_t bytes_received;

    while ((bytes_received = recv(sockfd, response, sizeof(response) -1, 0)) > 0) {
        response[bytes_received] = '\0';
        printf("%s", response);
    }

    if (bytes_received == -1) {
        print_error("recv");
    }

    close(sockfd);

    return 0;
}