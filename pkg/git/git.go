package git

import (
	"os/exec"
	"strings"
)

// GetLatestCommit returns a shorten version of latest commit's hash.
func GetLatestCommit() string {
	cmd := exec.Command("git", "log", "-n", "1", "--pretty=format:%h")

	output, _ := cmd.Output()

	return strings.TrimSpace(string(output))
}
