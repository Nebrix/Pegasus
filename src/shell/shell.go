package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"pegasus/src/tools"
	"pegasus/src/util"
)

type Shell struct {
	UserInfo UserInfo
}

type UserInfo struct {
	Username  string
	Directory string
	Hostname  string
	Home      string
}

func NewShell() *Shell {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting current user: %v\n", err)
		os.Exit(1)
	}

	currentWorkingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		os.Exit(1)
	}

	return &Shell{
		UserInfo: UserInfo{
			Username:  currentUser.Username,
			Directory: currentWorkingDir,
			Hostname:  hostname,
			Home:      currentUser.HomeDir,
		},
	}
}

func (s *Shell) Start() {
	s.commandLine()
}

func (s *Shell) commandLine() {
	reader := bufio.NewReader(os.Stdin)
	util.Ascii()

	for {
		currentDir := filepath.Base(s.UserInfo.Directory)
		fmt.Printf("[%s@%s %s]$ ", s.UserInfo.Username, s.UserInfo.Hostname, currentDir)
		userInput, err := getUserInput(reader)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		args := strings.Fields(userInput)

		if len(args) == 0 {
			continue
		}

		command := args[0]
		arguments := args[1:]

		switch command {
		case "exit":
			fmt.Println("Exiting the shell.")
			return
		case "cd":
			s.changeDirectory(arguments)
		case "ls":
			s.listDirectory()
		case "ping":
			s.ping(arguments)
		case "whois":
			s.whois(arguments)
		case "hash":
			s.hash(arguments)
		case "ip":
			s.showIP()
		case "dig":
			s.getDNSRecords(arguments)
		case "lookup":
			s.getIPInfo(arguments)
		case "clear":
			s.clearScreen()
		case "subnet":
			s.showSubnet(arguments)
		case "help":
			s.help()
		default:
			fmt.Printf("Command not found: %s\n", command)
		}
	}
}

func getUserInput(reader *bufio.Reader) (string, error) {
	userInput, err := reader.ReadString('\n')
	return strings.TrimSpace(userInput), err
}

func (s *Shell) clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (s *Shell) changeDirectory(arguments []string) {
	if len(arguments) != 1 {
		fmt.Println("Invalid usage of the cd command")
		return
	}

	newDir := arguments[0]
	if newDir == ".." {
		newDir = filepath.Dir(s.UserInfo.Directory)
	}

	err := os.Chdir(newDir)
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		return
	}

	s.UserInfo.Directory = newDir
}

func (s *Shell) listDirectory() {
	files, err := OSReadDir(s.UserInfo.Directory)
	if err != nil {
		fmt.Printf("Error listing directory: %v\n", err)
		return
	}

	for _, file := range files {
		fmt.Println(file)
	}
}

func (s *Shell) ping(arguments []string) {
	if len(arguments) < 2 {
		fmt.Println("Usage: ping <host> <count>")
		return
	}

	host := arguments[0]
	count, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Error parsing count:", err)
		return
	}

	tools.Ping(host, count, 2*time.Second)
}

func (s *Shell) whois(arguments []string) {
	if len(arguments) < 1 {
		fmt.Println("Usage: whois <host>")
		return
	}

	tools.Getwhois(arguments[0])
}

func (s *Shell) hash(arguments []string) {
	if len(arguments) < 2 {
		fmt.Println("Usage: hash <method> <data>")
		return
	}

	method := arguments[0]
	data := arguments[1]

	switch method {
	case "md5":
		tools.HashMD5(data)
	case "sha1":
		tools.HashSha1(data)
	case "sha256":
		tools.HashSha256(data)
	case "sha512":
		tools.HashSha512(data)
	case "decode":
		fmt.Printf("Hash Algorithm: %s\n", tools.Decodehash(data))
	default:
		fmt.Println("Unsupported hash method:", method)
	}
}

func (s *Shell) showIP() {
	tools.ShowIP()
}

func (s *Shell) showSubnet(arguments []string) {
	if len(arguments) < 2 {
		fmt.Println("Usage: subnet <host> <CIDR>")
		return
	}
	host := arguments[0]
	cidrStr := arguments[1]

	cidr, err := strconv.Atoi(cidrStr)
	if err != nil {
		fmt.Println("Invalid CIDR:", err)
		return
	}

	tools.ShowCalculation(host, cidr)
}

func (s *Shell) getDNSRecords(arguments []string) {
	if len(arguments) < 1 {
		fmt.Println("Usage: dig <hostname>")
		return
	}

	tools.GetDNSRecords(arguments[0])
}

func (s *Shell) getIPInfo(arguments []string) {
	if len(arguments) < 1 {
		fmt.Println("Usage: lookup <ipaddress>")
		return
	}

	tools.GetIpInfo(arguments[0])
}

func OSReadDir(root string) ([]string, error) {
	var files []string
	dirEntries, err := os.ReadDir(root)
	if err != nil {
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

func (s *Shell) help() {
	fmt.Println("ICMP Ping (ping): Send ICMP echo requests to check if a host is up.")
	fmt.Println("DNS Enumeration (dig): Perform DNS enumeration on a domain to gather information.")
	fmt.Println("WHOIS Lookup (whois): Retrieve WHOIS information for a domain.")
	fmt.Println("IP Lookup (lookup): Retrieve basic information about an IP address.")
	fmt.Println("Hash (hash): Generate a hash value.")
	fmt.Println("Ip (ip): gets local and public IP address for currently connected network.")
	fmt.Println("Subnet Calculator (subnet): Calculate subnet details and IP ranges.")
}
