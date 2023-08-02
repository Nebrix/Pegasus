import ipaddress

def subnet_calculator(ip_address, cidr):
    try:
        # Parse the IP address and CIDR notation
        ip = ipaddress.IPv4Address(ip_address)
        subnet = ipaddress.IPv4Network(f"{ip}/{cidr}", strict=False)

        # Calculate the network and broadcast addresses
        network_address = subnet.network_address
        broadcast_address = subnet.broadcast_address

        # Calculate the number of available addresses in the subnet
        host_count = len(list(subnet.hosts()))

        # Calculate the number of subnets
        subnet_count = 2**(32 - subnet.prefixlen)

        # Print subnet details
        print("Subnet Details:")
        print(f"IP address: {ip}")
        print(f"CIDR: /{cidr}")
        print(f"Network address: {network_address}")
        print(f"Broadcast address: {broadcast_address}")
        print(f"Number of available addresses: {host_count}")
        print(f"Number of subnets: {subnet_count}")

    except ValueError as e:
        print("Invalid IP address or CIDR notation:", e)

if __name__ == "__main__":
    ip_address = input("Enter IP address: ")
    cidr = int(input("Enter CIDR: "))
    subnet_calculator(ip_address, cidr)
