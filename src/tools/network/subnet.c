#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <inttypes.h>

void subnetCalculator(const char *ipAddress, int cidr) {
    struct in_addr ip;
    if (inet_pton(AF_INET, ipAddress, &ip) != 1) {
        printf("Invalid IP address\n");
        return;
    }

    uint32_t ipInt = ntohl(ip.s_addr);
    int ones = cidr, bits = 32;
    uint32_t mask = (1U << (bits - ones)) - 1;

    uint32_t networkInt = ipInt & ~mask;
    uint32_t broadcastInt = ipInt | mask;

    struct in_addr networkIP, broadcastIP;
    networkIP.s_addr = htonl(networkInt);
    broadcastIP.s_addr = htonl(broadcastInt);

    uint32_t availableAddresses = broadcastInt - networkInt + 1;
    uint32_t subnetCount = 1U << (bits - cidr);

    printf("Subnet Details:\n");
    printf("IP address: %s\n", ipAddress);
    printf("CIDR: /%d\n", cidr);
    printf("Network address: %s\n", inet_ntoa(networkIP));
    printf("Broadcast address: %s\n", inet_ntoa(broadcastIP));
    printf("Number of available addresses: %" PRIu32 "\n", availableAddresses);
    printf("Number of subnets: %" PRIu32 "\n", subnetCount);
}

int main(int argc, char *argv[]) {
    if (argc != 3) {
        return 1;
    }

    const char *ipAddress = argv[1];
    int cidr = atoi(argv[2]);

    if (cidr < 0 || cidr > 32) {
        printf("Invalid CIDR\n");
        return 1;
    }

    subnetCalculator(ipAddress, cidr);
    return 0;
}
