#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <unistd.h>

// Function to create an HTTP GET request
char *createGetRequest(const char *host, const char *path) {
    char *request = (char *)malloc(1024);
    if (request == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        exit(1);
    }

    snprintf(request, 1024, "GET %s HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", path, host);

    return request;
}

// Function to check if a directory exists on a remote server
int checkDirExists(const char *host, const char *path) {
    int sockfd;
    struct sockaddr_in server_addr;
    struct hostent *server;
    char *request;
    char buffer[1024];

    // Create a socket
    sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd < 0) {
        perror("Error creating socket");
        return 0;
    }

    server = gethostbyname(host);
    if (server == NULL) {
        fprintf(stderr, "Error: Host not found\n");
        return 0;
    }

    // Initialize server address structure
    memset(&server_addr, 0, sizeof(server_addr));
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(80);  // Default HTTP port
    memcpy(&server_addr.sin_addr.s_addr, server->h_addr, server->h_length);

    // Connect to the server
    if (connect(sockfd, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
        perror("Error connecting to server");
        close(sockfd);
        return 0;
    }

    // Create and send an HTTP GET request
    request = createGetRequest(host, path);
    send(sockfd, request, strlen(request), 0);

    // Read the response
    int bytes_received;
    while ((bytes_received = recv(sockfd, buffer, sizeof(buffer), 0)) > 0) {
        // Check if the response contains a successful status code (e.g., 200 OK)
        if (strstr(buffer, "200 OK") != NULL) {
            free(request);
            close(sockfd);
            return 1;
        }
    }

    free(request);
    close(sockfd);
    return 0;
}

int main(int argc, char *argv[]) {
    if (argc < 3) {
        fprintf(stderr, "Usage: %s <host> <wordlist_file>\n", argv[0]);
        return 1;
    }

    const char *host = argv[1];
    const char *wordlist_file = argv[2];

    FILE *file = fopen(wordlist_file, "r");
    if (file == NULL) {
        fprintf(stderr, "Error opening wordlist file\n");
        return 1;
    }

    char line[1024];

    printf("Directory buster started...\n");

    while (fgets(line, sizeof(line), file) != NULL) {
        // Remove newline character if present
        line[strcspn(line, "\r\n")] = 0;

        if (checkDirExists(host, line)) {
            printf("[GET] Directory found: %s/%s\n", host, line);
        }
    }

    fclose(file);

    return 0;
}
