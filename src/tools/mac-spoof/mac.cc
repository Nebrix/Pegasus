#include <iostream>
#include <iomanip>
#include <cstdlib>
#include <ctime>
#include <cstring>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <net/if.h>

void disableNetwork(char *netInterface) {
    char command[100];
    sprintf(command, "sudo ifconfig %s down", netInterface);

    int result = system(command);
}

void upNetwork(char *netInterface) {
    char command[100];
    sprintf(command, "sudo ifconfig %s up", netInterface);

    int result = system(command);
}

// Function to generate a random MAC address
void generateRandomMac(unsigned char* mac) {
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
        std::cerr << "Failed to create socket." << std::endl;
        return 1;
    }

    char interfaceName[IFNAMSIZ];
    std::cout << "Enter the interface name: ";
    std::cin >> interfaceName;
    disableNetwork(interfaceName);

    // Generate random MAC address
    unsigned char newMacBytes[6];
    generateRandomMac(newMacBytes);

    std::cout << "Generated MAC Address: ";
    for (int i = 0; i < 6; ++i) {
        std::cout << std::hex << std::setw(2) << std::setfill('0') << static_cast<int>(newMacBytes[i]);
        if (i < 5) {
            std::cout << ":";
        }
    }
    std::cout << std::dec << std::endl;

    // Set new MAC address
    struct ifreq ifr;
    memset(&ifr, 0, sizeof(ifr));
    strncpy(ifr.ifr_name, interfaceName, IFNAMSIZ - 1);
    memcpy(ifr.ifr_hwaddr.sa_data, newMacBytes, 6);

    if (setuid(0) != 0) {
        std::cerr << "Failed to set effective user ID to root." << std::endl;
        close(sock);
        upNetwork(interfaceName);
        return 1;
    }

    if (ioctl(sock, SIOCSIFHWADDR, &ifr) == -1) {
        std::cerr << "Failed to set new MAC address." << std::endl;
        close(sock);
        upNetwork(interfaceName);
        return 1;
    }

    close(sock);

    std::cout << "MAC address changed successfully." << std::endl;
    upNetwork(interfaceName);

    return 0;
}
