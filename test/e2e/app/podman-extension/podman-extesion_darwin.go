package podmanextension

import (
	"fmt"
	"github.com/adrianriobo/goax/pkg/util/delay"
	"os/exec"
)

const (
	PODMAN_INSTALLER_CONTINUE = "Continue"
	PODMAN_INSTALLER_AGREE    = "Agree"
	PODMAN_INSTALLER_INSTALL  = "Install"
	PODMAN_INSTALLER_CLOSE    = "Close"
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
	delay.Delay(delay.MEDIUM)
	pInstaller, err := autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pInstaller.Click(PODMAN_INSTALLER_CONTINUE, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	delay.Delay(delay.SMALL)

	pInstaller, err = pInstaller.Reload()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pInstaller.Click(PODMAN_INSTALLER_CONTINUE, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	delay.Delay(delay.SMALL)
	pInstaller, err = pInstaller.Reload()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pInstaller.Click(PODMAN_INSTALLER_AGREE, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	delay.Delay(delay.SMALL)
	pInstaller, err = pInstaller.Reload()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pInstaller.Click(PODMAN_INSTALLER_INSTALL, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	delay.Delay(delay.SMALL)
	pdInstallerSec := autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pdInstallerSec.SetValueOnFocus(userPassword)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	delay.Delay(delay.LONG)
	pInstaller, err = pInstaller.Reload()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pInstaller.Click(PODMAN_INSTALLER_CLOSE, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	return nil
}
