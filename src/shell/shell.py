import os
import subprocess
from colorama import Fore
from upper.ascii.ascii import ascii
from upper.coreUtil.ls import list_command

def help():
    print("Port Scanner (scanner): Scan for open ports on a target system.")
    print("ICMP Ping (ping): Send ICMP echo requests to check if a host is up.")
    print("DNS Enumeration (dns): Perform DNS enumeration on a domain to gather information.")
    print("WHOIS Lookup (whois): Retrieve WHOIS information for a domain.")
    print("Packet Sniffer (sniffer): Capture and analyze network packets.")
    print("Subnet Calculator (subnet): Calculate subnet details and IP ranges.")
    print("IP Lookup (lookup): Retrieve basic information about an IP address.")
    print("Hash Ident (hashid): Identify the type of hash.")
    print("Hash (hash): Generate a hash value.")
    print("Pegasus Edit (edit): Run an inline text editor.")
    print("Traceroute (route): Trace the route packets take to reach a destination host.")
    print("Get Ip (getip): gets local and public IP address for currently connected network.")
    print("MAC Address Spoofing (macspoof): Change the MAC address of a network interface to bypass network restrictions or enhance privacy.")

COMMANDS = {
    'ls': list_command,
    'ping': './dist/ping',
    'route': './dist/route',
    'getip': './dist/ip',
    'lookup': './dist/lookup',
    'whois': './dist/whois',
    'hashid': './dist/hashid',
    'hash': './dist/hash',
    'edit': './dist/editor',
    'dns': './dist/dns',
    'subnet': './dist/subnet',
    'macspoof': './dist/spoofmac',
    'dirb': './dist/dirb',
    'sniffer': './dist/sniffer'
}

def executeCommand(command, args):
    if command == 'ls':
        COMMANDS[command]()
    elif command in COMMANDS:
        if len(args) == 1:
            subprocess.run([COMMANDS[command], args[0]])
        elif len(args) == 2 and command == 'hash':
            subprocess.run([COMMANDS[command], args[0], args[1]])
        elif len(args) == 2 and command == 'subnet':
            subprocess.run([COMMANDS[command], args[0], args[1]])
        elif len(args) == 2 and command == 'dirb':
            subprocess.run([COMMANDS[command], args[0], args[1]])  
        else:
            print(f"Invalid arguments for {command} command.")
    else:
        os.system(command)

def init_shell():
    ascii()
    while True:
        userInput = input(f"{Fore.BLUE}{os.path.basename(os.getcwd())}{Fore.GREEN}➜ {Fore.WHITE}")

        if userInput.lower() == 'exit':
            break
        elif userInput.lower() == 'help':
            help()
        elif userInput.lower() == 'reload':
            ascii()

        parts = userInput.split()
        command = parts[0].lower()
        args = parts[1:]

        executeCommand(command, args)

if __name__ == '__main__':
    init_shell()
