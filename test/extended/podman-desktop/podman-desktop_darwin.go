package podmandesktop

import (
	"os/exec"
)

func cleanup() error {
	cmd := exec.Command("/bin/sh", "-c", "rm -rf ${HOME}/.local/share/containers/podman-desktop")
	return cmd.Start()
}
