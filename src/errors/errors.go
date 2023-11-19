package errors

import (
	"fmt"
	"shell/src/colors"
)

func HandleErr(msg string, err error) bool {
	if err != nil {
		fmt.Printf(colors.ERROR+"%v: %v\n"+colors.RESET, msg, err)
		return true
	}
	return false
}

func HandleWarn(msg string, p string) bool {
	fmt.Printf(colors.WARNING+"%v: %v\n"+colors.RESET, msg, p)
	return true
}
