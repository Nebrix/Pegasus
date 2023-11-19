package colors

import (
	"os"
	"runtime"
)

var (
	RESET   = "\033[0m"
	ERROR   = "\033[31m"
	SUCCESS = "\033[32m"
	WARNING = "\033[33m"
)

func init() {
	if runtime.GOOS == "windows" {
		if os.Stdout == nil || os.Stdout == os.Stderr {
			RESET = ""
			ERROR = ""
			SUCCESS = ""
			WARNING = ""
		}
	}
}
