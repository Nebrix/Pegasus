package helper

import (
	"fmt"
	"os/exec"
	"strings"
)

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

func Help() {
	fmt.Println("Available Commands:")
	fmt.Println("  - ICMP Ping (ping):\t\tSend ICMP echo requests to check if a host is up.")
	fmt.Println("  - DNS Enumeration (dig):\tPerform DNS enumeration on a domain to gather information.")
	fmt.Println("  - WHOIS Lookup (whois):\tRetrieve WHOIS information for a domain.")
	fmt.Println("  - IP Lookup (lookup):\t\tRetrieve basic information about an IP address.")
	fmt.Println("  - Hash (hash):\t\tGenerate a hash value.")
	fmt.Println("  - IP (ip):\t\t\tGet local and public IP addresses for the currently connected network.")
	fmt.Println("  - Subnet Calculator (subnet):\tCalculate subnet details and IP ranges.")
}
