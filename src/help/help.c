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