package main

import (
	"flag"
	"fmt"
	shell "shell/src"
)

func main() {
	styleFlag := flag.String("style", "default", "Specify the shell style")
	flag.Parse()

	config := shell.ShellConfig{PromptStyle: *styleFlag}
	initializeShell, err := shell.NewShell(config)
	if err != nil {
		fmt.Printf("Error initializing shell: %v\n", err)
		return
	}

	initializeShell.Start()
}
