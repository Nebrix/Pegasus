import time
import logging

from scapy.all import sniff

def packet_handler(packet):
    # Replace this with the packet processing logic you want
    print(packet.summary())

def get_default_network_device():
    from scapy.arch import get_if_list

    # Find all available network devices
    devices = get_if_list()

    # Return the name of the first network device (the default one)
    if devices:
        return devices[0]

    raise ValueError("No network devices found")

def main():
    # Get the default network device
    default_device = get_default_network_device()

    print("Starting packet sniffer on device:", default_device)

    # Start sniffing packets on the default network device
    try:
        sniff(iface=default_device, prn=packet_handler)
    except KeyboardInterrupt:
        print("Packet sniffer stopped.")

if __name__ == "__main__":
    logging.getLogger("scapy.runtime").setLevel(logging.ERROR)
    main()
