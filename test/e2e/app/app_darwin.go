package app

import (
	"os/exec"
)

func cleanup() error {
	// rmSharePodmanDesktop := "rm -rf $HOME/.local/share/containers/podman-desktop"
	// cmd = exec.Command(rmSharePodmanDesktop...)
	cmd := exec.Command("rm", "-rf", "${HOME}/.local/share/containers/podman-desktop")
	return cmd.Start()
}
