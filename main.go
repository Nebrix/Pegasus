package main

import (
	"flag"
	"fmt"
	shell "shell/src"
)

func main() {
	styleFlag := flag.String("style", "default", "Specify the shell style")
	flag.Parse()

	var initializeShell *shell.Shell
	switch *styleFlag {
	case "default":
		initializeShell = shell.NewShell(shell.ShellDefaults)
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
	default:
		fmt.Println("Invalid shell style. Available styles: default, windows")
		return
	}

	initializeShell.Start()
}
