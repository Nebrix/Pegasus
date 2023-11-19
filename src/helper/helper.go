package helper

import (
	"os/exec"
	"strings"
)

func ExtractUsername(p string) string {
	usernameParts := strings.Split(p, "\\")
	if len(usernameParts) > 1 {
		return usernameParts[1]
	}
	return p
}

func GetGitBranch(directory string) string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = directory

	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(output))
}
