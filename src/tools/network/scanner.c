#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <pthread.h>

#define NUM_THREADS 10 // You can adjust the number of threads

void *scan_port(void *data) {
    char *target_ip = (char *)data;

    for (int port = 1; port <= 65535; port++) {
        int sockfd = socket(AF_INET, SOCK_STREAM, 0);
        if (sockfd == -1) {
            perror("socket");
            exit(1);
        }

        struct sockaddr_in target_addr;
        target_addr.sin_family = AF_INET;
        target_addr.sin_port = htons(port);
        if (inet_pton(AF_INET, target_ip, &(target_addr.sin_addr)) <= 0) {
            perror("inet_pton");
            close(sockfd);
            continue;
        }

        if (connect(sockfd, (struct sockaddr *)&target_addr, sizeof(target_addr)) == 0) {
            printf("Port %d is open\n", port);
        }

        close(sockfd);
    }

    pthread_exit(NULL);
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        printf("Usage: %s <target IP>\n", argv[0]);
        exit(1);
    }

    char *target_ip = argv[1];
    pthread_t threads[NUM_THREADS];

    for (int i = 0; i < NUM_THREADS; i++) {
        if (pthread_create(&threads[i], NULL, scan_port, target_ip) != 0) {
            perror("pthread_create");
            exit(1);
        }
    }

    for (int i = 0; i < NUM_THREADS; i++) {
        pthread_join(threads[i], NULL);
    }

    return 0;
}
