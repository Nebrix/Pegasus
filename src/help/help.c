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

void Help() {
    puts("Port Scanner (scanner): Scan for open ports on a target system.");
    puts("ICMP Ping (ping): Send ICMP echo requests to check if a host is up.");
    puts("DNS Enumeration (dnslookup): Perform DNS enumeration on a domain to gather information.");
    puts("WHOIS Lookup (whois): Retrieve WHOIS information for a domain.");
    puts("Packet Sniffer (sniffer): Capture and analyze network packets.");
    puts("Subnet Calculator (subnet): Calculate subnet details and IP ranges.");
}
