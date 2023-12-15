package helper

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

var (
	RESET   = "\033[0m"
	ERROR   = "\033[31m"
	WARNING = "\033[33m"
)

func init() {
	if runtime.GOOS == "windows" {
		if os.Stdout == nil || os.Stdout == os.Stderr {
			RESET = ""
			ERROR = ""
			WARNING = ""
		}
	}
}

func ExtractUsername(p string) string {
	usernameParts := strings.Split(p, "\\")
	if len(usernameParts) > 1 {
		return usernameParts[1]
	}
	return p
}

func GetGitBranch(directory string) string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = directory

	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(output))
}

func MainHelp() {
	fmt.Println("Available Commands:")
	printCommandHelp("ping", "ICMP Ping", "Send ICMP echo requests to check the reachability of a host and measure round-trip times.")
	printCommandHelp("hash", "Hash", "Generate a cryptographic hash value for a given input.")
	printCommandHelp("subnet", "Subnet Calculator", "Calculate subnet details, including network and broadcast addresses, and IP ranges.")
	printCommandHelp("sniff", "Packet Sniffer", "Capture and analyze network packets on a specified interface.")
	printCommandHelp("traceroute", "Traceroute", "Reveal the network path and measure transit times of packets to a destination IP address.")
	printCommandHelp("scan", "Port Scanner", "Scan for open ports on a specified IP address or domain.")
}

func OsintHelp() {
	fmt.Println("Available Commands:")
	printCommandHelp("dig", "DNS Enumeration", "Perform DNS enumeration on a domain to gather information about its DNS records.")
	printCommandHelp("whois", "WHOIS Lookup", "Retrieve detailed registration information for a domain, including contact details.")
	printCommandHelp("lookup", "IP Lookup", "Retrieve basic information about an IP address, such as its geolocation and ISP.")
	printCommandHelp("webheader", "Website Header", "Retrieve basic header information via a HTTP web request")
	printCommandHelp("ip", "IP Addresses", "Display local and public IP addresses for the currently connected network.")
}

func printCommandHelp(command, title, description string) {
	fmt.Printf("  - %v (%v):\n\t%v\n", command, title, description)
}

func HandleErr(msg string, err error) bool {
	if err != nil {
		fmt.Printf(ERROR+"%v: %v\n"+RESET, msg, err)
		return true
	}
	return false
}

func HandleWarn(msg string, p string) bool {
	fmt.Printf(WARNING+"%v: %v\n"+RESET, msg, p)
	return true
}

func OSReadDir(root string) ([]string, error) {
	var files []string
	dirEntries, err := os.ReadDir(root)
	if HandleErr("", err) {
		return nil, err
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			files = append(files, dirEntry.Name()+"/")
		} else {
			files = append(files, dirEntry.Name())
		}
	}

	sort.Strings(files)

	return files, nil
}
