#include <iostream>
#include <cstring>
#include <cerrno>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <ifaddrs.h>
#include <netdb.h>

void getPublicIP(void (*callback)(const char *)) {
    const char *hostname = "api64.ipify.org";
    const char *path = "/?format=json";

    struct hostent *host = gethostbyname(hostname);
    if (host == nullptr) {
        std::cerr << "Error resolving hostname: " << hstrerror(h_errno) << std::endl;
        callback("Unknown");
        return;
    }

    struct sockaddr_in server_addr;
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(80);
    server_addr.sin_addr = *((struct in_addr *)host->h_addr);

    int sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd == -1) {
        perror("Error creating socket");
        callback("Unknown");
        return;
    }

    if (connect(sockfd, reinterpret_cast<struct sockaddr *>(&server_addr), sizeof(server_addr)) == -1) {
        perror("Error connecting to server");
        close(sockfd);
        callback("Unknown");
        return;
    }

    char request[256];
    snprintf(request, sizeof(request), "GET %s HTTP/1.1\r\nHost: %s\r\n\r\n", path, hostname);

    if (send(sockfd, request, strlen(request), 0) == -1) {
        perror("Error sending request");
        close(sockfd);
        callback("Unknown");
        return;
    }

    char response[1024];
    if (recv(sockfd, response, sizeof(response), 0) == -1) {
        perror("Error receiving response");
        close(sockfd);
        callback("Unknown");
        return;
    }

    char *ipStart = strstr(response, "\"ip\":\"") + 6;
    char *ipEnd = strstr(ipStart, "\"");
    if (ipStart == nullptr || ipEnd == nullptr) {
        std::cerr << "Error parsing public IP" << std::endl;
        close(sockfd);
        callback("Unknown");
        return;
    }

    *ipEnd = '\0';
    callback(ipStart);

    close(sockfd);
}

std::string getInternalIP() {
    const char* google_dns_server = "8.8.8.8";
    int dns_port = 53;

    struct sockaddr_in serv;
    int sock = socket(AF_INET, SOCK_DGRAM, 0);

    // Socket could not be created
    if (sock < 0) {
        std::cerr << "Socket error" << std::endl;
        return "Unknown";
    }

    memset(&serv, 0, sizeof(serv));
    serv.sin_family = AF_INET;
    serv.sin_addr.s_addr = inet_addr(google_dns_server);
    serv.sin_port = htons(dns_port);

    int err = connect(sock, reinterpret_cast<const struct sockaddr*>(&serv), sizeof(serv));
    if (err < 0) {
        std::cerr << "Error number: " << errno
                  << ". Error message: " << strerror(errno) << std::endl;
        close(sock);
        return "Unknown";
    }

    struct sockaddr_in name;
    socklen_t namelen = sizeof(name);
    err = getsockname(sock, reinterpret_cast<struct sockaddr*>(&name), &namelen);

    char buffer[80];
    const char* p = inet_ntop(AF_INET, &name.sin_addr, buffer, sizeof(buffer));
    if (p != nullptr) {
        close(sock);
        return std::string(buffer);
    } else {
        std::cerr << "Error number: " << errno
                  << ". Error message: " << strerror(errno) << std::endl;
        close(sock);
        return "Unknown";
    }
}

void printIP(const char *ip) {
    std::cout << "IP: " << ip << std::endl;
}

void printIPs(const char *internalIP, const char *publicIP) {
    std::cout << "Internal IP: " << internalIP << std::endl;
    std::cout << "Public IP: " << publicIP << std::endl;
}

void mainCallback(const char *ip) {
    std::string internalIP = getInternalIP();
    printIPs(internalIP.c_str(), ip);
}

int main() {
    getPublicIP(mainCallback);
    return 0;
}
