#include <stdio.h>
#include "help.h"

void pingHelp() {
    puts("Usage: ping <dest IP>");
}

void scannerHelp() {
    puts("Usage: scanner <dest IP> <start port> <end port>");
}

void dnslookupHelp() {
    puts("Usage: dnslookup <URL>");
}

void whoisHelp() {
    puts("Usage: whois <IP || URL>");
}

void dirbHelp() {
    puts("Usage: dirb <URL>");
}
void hashHelp() {
    puts("Usage: hash <string> <alg>");
}

void hashIdentHelp() {
    puts("Usage: hashident <string>");
}

void Help() {
    puts("Port Scanner (scanner): Scan for open ports on a target system.");
    puts("ICMP Ping (ping): Send ICMP echo requests to check if a host is up.");
    puts("DNS Enumeration (dnslookup): Perform DNS enumeration on a domain to gather information.");
    puts("WHOIS Lookup (whois): Retrieve WHOIS information for a domain.");
    puts("Packet Sniffer (sniffer): Capture and analyze network packets.");
    puts("Subnet Calculator (subnet): Calculate subnet details and IP ranges.");
    puts("IP Lookup (iplookup): Retrieve basic information about an IP address.");
    puts("Hash Ident (hashident): Identify the type of hash.");
    puts("Hash (hash): Generate a hash value.");
    puts("Server (server): Create a server on localhost that can be connected to using telnet or ncat.");
    puts("Pegasus Edit (edit): Run an inline text editor.");
    puts("Traceroute (traceroute): Trace the route packets take to reach a destination host.");
    puts("Web Server (webserver): Run a simple web server for quick file sharing or testing purposes.");
    puts("Reverse Shell (revshell): Create a reverse shell listener to establish a network connection to a remote system.");
}
