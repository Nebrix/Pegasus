#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <netdb.h>

#define HOSTNAME "ipinfo.io"
#define PORT "80"

void print_error(const char *msg) {
    fprintf(stderr, "Error: %s\n", msg);
    exit(1);
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <IP Address>\n", argv[0]);
        exit(1);
    }

    char *ipAddress = argv[1];
    char path[256];
    snprintf(path, sizeof(path), "/%s/json", ipAddress);

    struct addrinfo hints, *res;
    int sockfd;

    memset(&hints, 0, sizeof(hints));
    hints.ai_family = AF_INET;
    hints.ai_socktype = SOCK_STREAM;

    if (getaddrinfo(HOSTNAME, PORT, &hints, &res) != 0) {
        print_error("getaddrinfo");
    }

    sockfd = socket(res->ai_family, res->ai_socktype, res->ai_protocol);
    if (sockfd == -1) {
        print_error("socket");
    }

    if (connect(sockfd, res->ai_addr, res->ai_addrlen) == -1) {
        print_error("connect");
    }

    char request[256];
    snprintf(request, sizeof(request), "GET %s HTTP/1.1\r\nHost: %s\r\nUser-Agent: C HTTP Client\r\n\r\n", path, HOSTNAME);

    if (send(sockfd, request, strlen(request), 0) == -1) {
        print_error("send");
    }

    char response[4096];
    ssize_t bytes_received;
    int header_end = 0; // Flag to track the end of HTTP headers

    while ((bytes_received = recv(sockfd, response, sizeof(response) - 1, 0)) > 0) {
        response[bytes_received] = '\0';
        printf("%s", response);

        // Check for the end of headers (empty line)
        if (strstr(response, "\r\n\r\n") != NULL) {
            header_end = 1;
            break;
        }
    }

    if (bytes_received == -1) {
        print_error("recv");
    }

    close(sockfd);
    freeaddrinfo(res);

    if (!header_end) {
        // Headers are incomplete or missing, handle the error.
        print_error("Incomplete or missing headers in the response.");
    }

    return 0;
}
