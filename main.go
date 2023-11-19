package main

import (
	"flag"
	shell "shell/src"
)

func main() {
	styleFlag := flag.String("style", "default", "Specify the shell style")
	flag.Parse()

	var initializeShell *shell.Shell
	switch *styleFlag {
	case "windows":
		windowsConfig := shell.ShellConfig{
			PromptStyle: "windows",
		}
		initializeShell = shell.NewShell(windowsConfig)
	case "zsh":
		zshConfig := shell.ShellConfig{
			PromptStyle: "zsh",
		}
		initializeShell = shell.NewShell(zshConfig)
	case "zsh-git":
		zshGitConfig := shell.ShellConfig{
			PromptStyle: "zsh-git",
		}
		initializeShell = shell.NewShell(zshGitConfig)
	case "root":
		rootConfig := shell.ShellConfig{
			PromptStyle: "root",
		}
		initializeShell = shell.NewShell(rootConfig)
	case "mac":
		macConfig := shell.ShellConfig{
			PromptStyle: "mac",
		}
		initializeShell = shell.NewShell(macConfig)
	case "hacker":
		hackerConfig := shell.ShellConfig{
			PromptStyle: "hacker",
		}
		initializeShell = shell.NewShell(hackerConfig)
	default:
		initializeShell = shell.NewShell(shell.ShellDefaults)
	}

	initializeShell.Start()
}
