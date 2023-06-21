package app

import (
	"os/exec"
)

func cleanup() error {
	rmSharePodmanDesktop := "rm -rf $HOME/.local/share/containers/podman-desktop"
	cmd = exec.Command(rmSharePodmanDesktop...)
	return cmd.Start()
}
