package main

import (
	"dedsec/src/shell"
)

func main() {
	initShell := shell.NewShell()
	initShell.Start()
}
