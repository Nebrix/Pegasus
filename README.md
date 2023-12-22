# Pegasus - OS-Based Hacking Shell

<p align="center">
    <img src="images/pegasus.png" alt="pegasus logo">
</p>

<!-- TOC -->
- [Introduction](#introduction)
- [Prompt Styles](#prompt-styles)
- [Docker Installation](#docker-installation)
    - [Local Install](#local-install)
    - [Docker Hub Install](#docker-hub-install)
- [Contribution](#contribution)
- [License](#license)
- [Disclaimer](#disclaimer)
- [Todo](#todo)
<!-- TOC -->

## Introduction

Pegasus is a powerful hacking shell designed for Unix-based operating systems. It provides various tools and functionalities that can be used for security testing and ethical hacking purposes. This tool is intended for educational and responsible use only. Please use it responsibly and with proper authorization.

## Prompt styles

- windows
- root
- zsh, zsh-git
- mac
- hacker

## Docker installation

### Local install
If you prefer to build the Docker image locally, execute the following commands:

`docker build -t pegasus .`

`docker run -it pegasus`

### Docker hub install (Recommend)
For a more straightforward installation, you can pull the Docker image from Docker Hub:

`docker pull nebrix/pegasus:4.5.2`

`docker run -it docker.io/nebrix/pegasus:4.5.2`

Using the Docker Hub image is the recommended and easier approach for most users.

## Todo

- [X] Whois- [Description: Retrieve detailed registration information for a domain, including contact details]
- [X] DNS - [Description: Perform DNS enumeration on a domain to gather information about its DNS records]
- [X] Hashing - [Description: Generate a cryptographic hash value for a given input]
- [X] IP/IP Information - [Description: Retrieve basic information about an IP address, such as its geolocation and ISP]
- [X] Subnet Calculator - [Description: Calculate subnet details, including network and broadcast addresses, and IP ranges]

- [X] Port Scanner - [Description: Scan for open ports on a specified IP address or domain]
- [X] Packet Sniffer - [Description: Capture and analyze network packets on a specified interface]
- [ ] Discover WiFi Networks - [Description: Discover networks]

- [X] Ping - [Description: Send ICMP echo requests to check the reachability of a host and measure round-trip times]
- [X] Traceroute - [Description: Reveal the network path and measure transit times of packets to a destination IP address]
- [X] Web Header - [Description: Retrieve basic header information via an HTTP web request]
- [X] IP Addresses - [Description: Display local and public IP addresses for the currently connected network]

- [X] Shell Prompt Styles - [Description: Customize the style of the shell prompt]
- [ ] Custom Prompt Styles - [Description: Create custom shell prompts]

- [X] Install Script - [Description: Installer for libpcap (Unix only)]

## Contribution

If you find any bugs or want to contribute to Pegasus, please feel free to open an issue or submit a pull request on the GitHub repository. We welcome your feedback and suggestions to make this tool even better.

## License

Pegasus is open-source software licensed under the [MIT License](https://github.com/Nebrix/Pegasus/blob/main/COPYING). You are free to use, modify, and distribute this software with proper attribution and in compliance with the license terms.

## Disclaimer

Pegasus is provided for educational and ethical hacking purposes only. The authors and contributors of Pegasus are not responsible for any misuse or illegal activities performed using this tool. Please use it responsibly and in compliance with the laws and regulations of your country.
