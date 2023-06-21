package podmanextension

import (
	"fmt"
	"os/exec"
)

func cleanup() error {
	// rmConfigContainers := "rm -rf $HOME/.config/containers"
	// cmd := exec.Command(rmConfigContainers...)
	cmd := exec.Command("rm", "-rf", "$HOME/.config/containers")
	err := cmd.Start()
	if err != nil {
		return err
	}
	// rmShareContainers := "rm -rf $HOME/.local/share/containers"
	// cmd = exec.Command(rmShareContainers...)
	cmd = exec.Command("rm", "-rf", "$HOME/.local/share/containers")
	err = cmd.Start()
	if err != nil {
		return err
	}
	// rmPodman := "sudo rm -rf /opt/podman"
	// cmd = exec.Command(rmPodman...)
	cmd = exec.Command("sudo", "rm", "-rf", "/opt/podman")
	return cmd.Start()
}

func installer(userPassword string) error {
	return fmt.Errorf("not implemented yet")
}
