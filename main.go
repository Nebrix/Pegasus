package main

import (
	"pegasus/src/shell"
)

func main() {
	initShell := shell.NewShell()
	initShell.Start()
}
