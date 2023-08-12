#include <stdio.h>
#include <stdlib.h>
#include <pcap.h>
#include <signal.h>
#include <string.h>
#include <unistd.h>

void packet_handler(u_char *user_data, const struct pcap_pkthdr *pkthdr, const u_char *packet) {
    // Replace this with the packet processing logic you want
    printf("Received a packet\n");
}

char* get_default_network_device() {
    pcap_if_t *alldevs;
    char errbuf[PCAP_ERRBUF_SIZE];

    if (pcap_findalldevs(&alldevs, errbuf) == -1) {
        fprintf(stderr, "Error finding devices: %s\n", errbuf);
        exit(1);
    }

    char *device = strdup(alldevs->name);
    pcap_freealldevs(alldevs);

    return device;
}

void signal_handler(int signo) {
    printf("Packet sniffer stopped.\n");
    exit(0);
}

int main() {
    signal(SIGINT, signal_handler);
    signal(SIGTERM, signal_handler);

    char *default_device = get_default_network_device();
    printf("Starting packet sniffer on device: %s\n", default_device);

    char errbuf[PCAP_ERRBUF_SIZE];
    pcap_t *handle = pcap_open_live(default_device, BUFSIZ, 1, 1000, errbuf);
    if (handle == NULL) {
        fprintf(stderr, "Could not open device %s: %s\n", default_device, errbuf);
        return 1;
    }

    struct pcap_pkthdr *header;
    const u_char *packet;

    while (1) {
        int res = pcap_next_ex(handle, &header, &packet);
        if (res == 1) {
            packet_handler(NULL, header, packet);
        } else if (res == 0) {
            // Timeout
            continue;
        } else if (res == -1 || res == -2) {
            // Error or EOF
            break;
        }
    }

    pcap_close(handle);
    free(default_device);

    return 0;
}
