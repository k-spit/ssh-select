package gnome

import (
	"fmt"
	"os"
	"os/exec"

	sshselect "github.com/iwittkau/ssh-select"
)

// NewSSHTerminalWindow opens an SSH connection in the current terminal window
func NewSSHTerminalWindow(server sshselect.Server) error {
	var cmd *exec.Cmd
	if server.Port == "" {
		cmd = exec.Command("ssh", fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	} else {
		cmd = exec.Command("ssh", "-p", server.Port, fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
