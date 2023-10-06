# Pegasus - Unix-Based Hacking Shell

<p align="center">
    <img src="images/pegasus.png" alt="pegasus logo">
</p>

## Introduction

Pegasus is a powerful hacking shell designed for Unix-based operating systems. It provides various tools and functionalities that can be used for security testing and ethical hacking purposes. This tool is intended for educational and responsible use only. Please use it responsibly and with proper authorization.

# Important 

You cannot run telnet inside of the shell or it will kill the shell (Still trouble shooting)

### Note

If you are using Windows, you would need to use a VM or a Linux subsystem (WSL) to run Pegasus.


### Note 

you need to install pcap for your distro

## Installation

To install Pegasus, follow these steps:

1. Clone the repository to your local machine:

`git clone https://github.com/Nebrix/Pegasus.git`

2. Change into the Pegasus directory:

`cd Pegasus`

3. Make the build script executable:

`chmod a+x build`

4. Run the build script:

`./build`

## Usage

Once Pegasus is successfully installed, you can run it by executing the `sudo python3 src/main.py` command in your terminal:

![Pegasus Terminal](https://github.com/Codezz-ops/Pegasus/assets/112660193/32d2fd19-b35d-469c-935c-34eb8f28d95c)

## Features

Pegasus comes with a variety of hacking and security testing tools, including:

- Port Scanner (scanner): Scan for open ports on a target system.
- ICMP Ping (ping): Send ICMP echo requests to check if a host is up.
- DNS Enumeration (dnslookup): Perform DNS enumeration on a domain to gather information.
- WHOIS Lookup (whois): Retrieve WHOIS information for a domain.
- Packet Sniffer (sniffer): Capture and analyze network packets.
- Subnet Calculator (subnet): Calculate subnet details and IP ranges.
- IP Lookup (iplookup): Retrieve basic information about an IP address.
- Hash Ident (hashident): Identify the type of hash.
- Hash (hash): Generate a hash value.
- Server (server): Create a server on localhost that can be connected to using telnet or ncat.
- Pegasus Edit (edit): Run an inline text editor.
- Traceroute (traceroute): Trace the route packets take to reach a destination host.
- Get Ip (getip): gets local and public IP address for currently connected network.
- MAC Address Spoofing (macspoof): Change the MAC address of a network interface to bypass network restrictions or enhance privacy.

## Contribution

If you find any bugs or want to contribute to Pegasus, please feel free to open an issue or submit a pull request on the GitHub repository. We welcome your feedback and suggestions to make this tool even better.

## License

Pegasus is open-source software licensed under the [MIT License](https://github.com/codezz-ops/pegasus/blob/main/LICENSE). You are free to use, modify, and distribute this software with proper attribution and in compliance with the license terms.

## Disclaimer

Pegasus is provided for educational and ethical hacking purposes only. The authors and contributors of Pegasus are not responsible for any misuse or illegal activities performed using this tool. Please use it responsibly and in compliance with the laws and regulations of your country.

This tool is intended for educational and responsible use only. The authors and contributors of Pegasus are not responsible for any misuse or illegal activities performed using this tool. Use it at your own risk and always seek proper authorization before using it on systems you do not own or control.
