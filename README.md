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

If you use the Experimental branch, expect things to not function correctly.

## Installation

To install Pegasus, follow these steps:

1. Clone the repository to your local machine:

`git clone https://github.com/Codezz-ops/Pegasus.git`

2. Change into the Pegasus directory:

`cd Pegasus`

3. Make the build & install script executable:

`chmod a+x build/build.sh`
`chmod a+x build/install`

4. Run the build & install script:

`./build/install`
`./build/build.sh`

## Usage

Once Pegasus is successfully installed, you can run it by executing the `pegasus` command in your terminal:

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
- Web Server (webserver): Run a simple web server for quick file sharing or testing purposes.
- Reverse Shell (revshell): Create a reverse shell listener to establish a network connection to a remote system.
- Get Ip (getip): gets local and public IP address for currently connected network.

## Future Updates

We are continuously improving Pegasus and working on exciting new features:

- Working on implemting an inline text editor :heavy_check_mark:
- Working on implemting an inline chat room for connection on same network connection :heavy_check_mark:

## Contribution

If you find any bugs or want to contribute to Pegasus, please feel free to open an issue or submit a pull request on the GitHub repository. We welcome your feedback and suggestions to make this tool even better.

## License

Pegasus is open-source software licensed under the [MIT License](https://github.com/codezz-ops/pegasus/blob/main/LICENSE). You are free to use, modify, and distribute this software with proper attribution and in compliance with the license terms.

## Disclaimer

Pegasus is provided for educational and ethical hacking purposes only. The authors and contributors of Pegasus are not responsible for any misuse or illegal activities performed using this tool. Please use it responsibly and in compliance with the laws and regulations of your country.

This tool is intended for educational and responsible use only. The authors and contributors of Pegasus are not responsible for any misuse or illegal activities performed using this tool. Use it at your own risk and always seek proper authorization before using it on systems you do not own or control.
