package podmanextension

import (
	"fmt"
	"os/exec"
)

func cleanup() error {
	cmd := exec.Command("/bin/sh", "-c", "rm -rf ${HOME}/.config/containers")
	err := cmd.Start()
	if err != nil {
		return err
	}
	cmd = exec.Command("/bin/sh", "-c", "rm -rf ${HOME}/.local/share/containers")
	err = cmd.Start()
	if err != nil {
		return err
	}
	cmd = exec.Command("sudo", "rm", "-rf", "/opt/podman")
	return cmd.Start()
}

func installer(userPassword string) error {
	return fmt.Errorf("not implemented yet")
}
