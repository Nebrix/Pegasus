package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"shell/src/cmd"
	"shell/src/helper"
	"strings"
)

const (
	WindowsPrompt = "windows"
	ZshPrompt     = "zsh"
	ZshGitPrompt  = "zsh-git"
	RootPrompt    = "root"
	MacPrompt     = "mac"
	HackerPrompt  = "hacker"
	DefaultPrompt = "default"

	EscapeWindows = "%v> \033[0m"
	EscapeZsh     = "\033[1;32m➜  \033[1;34m%v\033[0m "
	EscapeZshGit  = "\033[1;32m➜  \033[1;34m%v \033[31m(%v)\033[0m "
	EscapeRoot    = "%v# "
	EscapeMac     = "%v:%v$ "
	EscapeHacker  = "\033[32m%v=%v?\033[0m "
	EscapeDefault = "[%v@%v %v]$ "
)

type ShellConfig struct {
	PromptStyle string
}

var ShellDefaults = ShellConfig{
	PromptStyle: DefaultPrompt,
}

type Shell struct {
	UserInfo    UserInfo
	ShellConfig ShellConfig
	shellName   string
}

type UserInfo struct {
	Username  string
	Directory string
	Hostname  string
	Home      string
}

type OSINTShell struct {
	mainShell *Shell
}

func NewShell(config ShellConfig) (*Shell, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("error getting current user: %v", err)
	}

	currentWorkingDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current working directory: %v", err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("error getting hostname: %v", err)
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
		shellName:   "default",
	}, nil
}

func setPrompt(s *Shell) {
	switch strings.ToLower(s.ShellConfig.PromptStyle) {
	case WindowsPrompt:
		setEscapePrompt(EscapeWindows, s.UserInfo.Directory)
	case ZshPrompt:
		setEscapePrompt(EscapeZsh, filepath.Base(s.UserInfo.Directory))
	case ZshGitPrompt:
		gitBranch := helper.GetGitBranch(s.UserInfo.Directory)
		setEscapePrompt(EscapeZshGit, filepath.Base(s.UserInfo.Directory), gitBranch)
	case RootPrompt:
		setEscapePrompt(EscapeRoot, filepath.Base(s.UserInfo.Directory))
	case MacPrompt:
		setEscapePrompt(EscapeMac, s.UserInfo.Username, filepath.Base(s.UserInfo.Directory))
	case HackerPrompt:
		setEscapePrompt(EscapeHacker, s.UserInfo.Username, filepath.Base(s.UserInfo.Directory))
	default:
		setEscapePrompt(EscapeDefault, s.UserInfo.Username, s.UserInfo.Hostname, filepath.Base(s.UserInfo.Directory))
	}
}

func setEscapePrompt(escapeFormat string, args ...interface{}) {
	fmt.Printf(escapeFormat, args...)
}

func (s *Shell) setOSINTPrompt() {
	fmt.Printf("[%v %v]$ ", s.shellName, filepath.Base(s.UserInfo.Directory))
}

func (s *Shell) Start() {
	s.commandLine()
}

func (s *Shell) commandLine() {
	reader := bufio.NewReader(os.Stdin)
	helper.Ascii()

	for {
		setPrompt(s)

		userInput, err := reader.ReadString('\n')
		if handleInputError(err) {
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
		case "version":
			fmt.Printf("%v\n", helper.GetVersion())
		case "cls", "clear", "Clear-Host":
			clearScreen()
		case "cd", "Set-Location":
			s.changeDirectory(arguments)
		case "ls", "dir":
			s.listDirectory()
		case "ping":
			cmd.Ping(arguments)
		case "hash":
			cmd.Hash(arguments)
		case "subnet":
			cmd.ShowSubnet(arguments)
		case "sniff":
			cmd.ShowSnifferPackets(arguments)
		case "scan":
			cmd.PortScanner(arguments)
		case "traceroute":
			cmd.Traceroute(arguments)
		case "help", "man":
			helper.MainHelp()
		case "osint":
			s.shellName = "osint"
			osintShell := newOSINTShell(s)
			osintShell.start()
		default:
			helper.HandleWarn("Command not found", userInput)
		}
	}
}

func (osintShell *OSINTShell) start() {
	reader := bufio.NewReader(os.Stdin)
	for {
		osintShell.mainShell.setOSINTPrompt()

		userInput, err := reader.ReadString('\n')
		if handleInputError(err) {
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
			return
		case "version":
			fmt.Printf("%v\n", helper.GetVersion())
		case "cls", "clear", "Clear-Host":
			clearScreen()
		case "cd", "Set-Location":
			osintShell.mainShell.changeDirectory(arguments)
		case "ls", "dir":
			osintShell.mainShell.listDirectory()
		case "whois":
			cmd.Whois(arguments)
		case "ip":
			cmd.ShowIP()
		case "dig":
			cmd.GetDNSRecords(arguments)
		case "lookup":
			cmd.GetIPInfo(arguments)
		case "webheader":
			cmd.Webheader(arguments)
		case "phonedns":
			cmd.PhoneDNS()
		case "help", "man":
			helper.OsintHelp()
		default:
			helper.HandleWarn("Command not found", userInput)
		}
	}
}

func handleInputError(err error) bool {
	helper.HandleErr("Error reading input", err)
	return false
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func newOSINTShell(mainShell *Shell) *OSINTShell {
	return &OSINTShell{mainShell: mainShell}
}

func (s *Shell) changeDirectory(arguments []string) {
	newDir := arguments[0]
	if newDir == ".." {
		newDir = filepath.Dir(s.UserInfo.Directory)
	}

	if helper.HandleErr("Error changing directory", os.Chdir(newDir)) {
		return
	}

	s.UserInfo.Directory = newDir
}

func (s *Shell) listDirectory() {
	if files, err := helper.OSReadDir(s.UserInfo.Directory); !helper.HandleErr("Error listing directory", err) {
		fmt.Println(strings.Join(files, "\n"))
	}
}
