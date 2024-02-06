package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetLatestCommit returns a shorten version of latest commit's hash.
func GetLatestCommit() (string, error) {
	cmd := exec.Command("git", "log", "-n", "1", "--pretty=format:%h")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
