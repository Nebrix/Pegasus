#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <net/if.h>
#include <time.h>

void disableNetwork(char *netInterface) {
    char command[100];
    snprintf(command, sizeof(command), "sudo ifconfig %s down", netInterface);

    int result = system(command);
    if (result != 0) {
        perror("Failed to disable network");
        exit(1);
    }
}

void upNetwork(char *netInterface) {
    char command[100];
    snprintf(command, sizeof(command), "sudo ifconfig %s up", netInterface);

    int result = system(command);
    if (result != 0) {
        perror("Failed to enable network");
        exit(1);
    }
}

// Function to generate a random MAC address
void generateRandomMac(unsigned char *mac) {
    srand(time(NULL));
    for (int i = 0; i < 6; ++i) {
        mac[i] = rand() & 0xFF;
    }
    // Set the locally administered bit and clear the multicast bit
    mac[0] |= 0x02;
    mac[0] &= 0xFE;
}

int main() {
    int sock = socket(AF_INET, SOCK_DGRAM, 0);
    if (sock == -1) {
        perror("Failed to create socket");
        exit(1);
    }

    char interfaceName[IFNAMSIZ];
    printf("Enter the interface name: ");
    scanf("%s", interfaceName);
    disableNetwork(interfaceName);

    // Generate random MAC address
    unsigned char newMacBytes[6];
    generateRandomMac(newMacBytes);

    printf("Generated MAC Address: ");
    for (int i = 0; i < 6; ++i) {
        printf("%02X", newMacBytes[i]);
        if (i < 5) {
            printf(":");
        }
    }
    printf("\n");

    // Set new MAC address
    struct ifreq ifr;
    memset(&ifr, 0, sizeof(ifr));
    strncpy(ifr.ifr_name, interfaceName, IFNAMSIZ - 1);
    memcpy(ifr.ifr_hwaddr.sa_data, newMacBytes, 6);

    if (setuid(0) != 0) {
        perror("Failed to set effective user ID to root");
        close(sock);
        upNetwork(interfaceName);
        exit(1);
    }

    if (ioctl(sock, SIOCSIFHWADDR, &ifr) == -1) {
        perror("Failed to set new MAC address");
        close(sock);
        upNetwork(interfaceName);
        exit(1);
    }

    close(sock);

    printf("MAC address changed successfully.\n");
    upNetwork(interfaceName);

    return 0;
}
