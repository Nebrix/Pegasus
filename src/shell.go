package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"shell/src/errors"
	"shell/src/helper"
	"shell/src/tools"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ShellConfig struct {
	PromptStyle string
}

var ShellDefaults = ShellConfig{
	PromptStyle: "default",
}

type Shell struct {
	UserInfo    UserInfo
	ShellConfig ShellConfig
}

type UserInfo struct {
	Username  string
	Directory string
	Hostname  string
	Home      string
}

func NewShell(config ShellConfig) *Shell {
	currentUser, err := user.Current()
	if errors.HandleErr("Error getting current user", err) {
		os.Exit(1)
	}

	currentWorkingDir, err := os.Getwd()
	if errors.HandleErr("Error getting current working directory", err) {
		os.Exit(1)
	}

	hostname, err := os.Hostname()
	if errors.HandleErr("Error getting hostname", err) {
		os.Exit(1)
	}

	username := helper.ExtractUsername(currentUser.Username)

	return &Shell{
		UserInfo: UserInfo{
			Username:  username,
			Directory: currentWorkingDir,
			Hostname:  hostname,
			Home:      currentUser.HomeDir,
		},
		ShellConfig: config,
	}
}

func (s *Shell) setPrompt() {
	switch s.ShellConfig.PromptStyle {
	case "windows":
		currentDrive := filepath.VolumeName(s.UserInfo.Directory)
		relativePath := strings.TrimPrefix(s.UserInfo.Directory, currentDrive)
		relativePath = strings.TrimPrefix(relativePath, string(os.PathSeparator))
		relativePath = strings.ReplaceAll(relativePath, "/", "\\")
		fmt.Printf("%v\\%v> ", currentDrive, relativePath)
	case "zsh":
		currentDir := filepath.Base(s.UserInfo.Directory)
		fmt.Printf("\033[1;32m➜  \033[1;34m%v\033[0m ", currentDir)
	case "zsh-git":
		gitBranch := helper.GetGitBranch(s.UserInfo.Directory)
		currentDir := filepath.Base(s.UserInfo.Directory)
		fmt.Printf("\033[1;32m➜  \033[1;34m%v \033[31m(%v)\033[0m ", currentDir, gitBranch)
	case "root":
		currentDir := filepath.Base(s.UserInfo.Directory)
		fmt.Printf("%v# ", currentDir)
	case "mac":
		currentDir := filepath.Base(s.UserInfo.Directory)
		fmt.Printf("%v:%v$ ", s.UserInfo.Username, currentDir)
	default:
		currentDir := filepath.Base(s.UserInfo.Directory)
		fmt.Printf("[%v@%v %v]$ ", s.UserInfo.Username, s.UserInfo.Hostname, currentDir)
	}
}

func (s *Shell) Start() {
	s.commandLine()
}

func (s *Shell) commandLine() {
	reader := bufio.NewReader(os.Stdin)
	Ascii()

	for {
		s.setPrompt()

		userInput, err := reader.ReadString('\n')
		if errors.HandleErr("Error reading input", err) {
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
			os.Exit(0)
		case "cls", "clear", "Clear-Host":
			fmt.Print("\033[H\033[2J")
		case "cd", "Set-Location":
			s.changeDirectory(arguments)
		case "ls", "dir":
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
		case "subnet":
			s.showSubnet(arguments)
		case "help", "man":
			helper.Help()
		default:
			errors.HandleWarn("Command not found", userInput)
		}
	}
}

func (s *Shell) changeDirectory(arguments []string) {
	newDir := arguments[0]
	if newDir == ".." {
		newDir = filepath.Dir(s.UserInfo.Directory)
	}

	err := os.Chdir(newDir)
	if errors.HandleErr("Error changing directory", err) {
		return
	}

	s.UserInfo.Directory = newDir
}

func (s *Shell) listDirectory() {
	files, err := OSReadDir(s.UserInfo.Directory)
	if errors.HandleErr("Error listing directory", err) {
		return
	}

	for _, file := range files {
		fmt.Println(file)
	}
}

func OSReadDir(root string) ([]string, error) {
	var files []string
	dirEntries, err := os.ReadDir(root)
	if errors.HandleErr("", err) {
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

func (s *Shell) ping(arguments []string) {
	host := arguments[0]
	count, err := strconv.Atoi(arguments[1])
	if errors.HandleErr("Error parsing count", err) {
		return
	}

	tools.Ping(host, count, 2*time.Second)
}

func (s *Shell) whois(arguments []string) {
	tools.Getwhois(arguments[0])
}

func (s *Shell) hash(arguments []string) {
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
		errors.HandleWarn("Unsupported hash method", method)
	}
}

func (s *Shell) showIP() {
	tools.ShowIP()
}

func (s *Shell) showSubnet(arguments []string) {
	host := arguments[0]
	cidrStr := arguments[1]

	cidr, err := strconv.Atoi(cidrStr)
	if errors.HandleErr("Invalid CIDR", err) {
		return
	}

	tools.ShowCalculation(host, cidr)
}

func (s *Shell) getDNSRecords(arguments []string) {
	tools.GetDNSRecords(arguments[0])
}

func (s *Shell) getIPInfo(arguments []string) {
	tools.GetIpInfo(arguments[0])
}
