package podmanextension

import (
	"fmt"
	autoApp "github.com/adrianriobo/goax/pkg/goax/app"
	"github.com/adrianriobo/goax/pkg/util/delay"
	"github.com/adrianriobo/goax/pkg/util/logging"
	"os/exec"
)

const (
	PODMAN_INSTALLER_CONTINUE = "Continue"
	PODMAN_INSTALLER_AGREE    = "Agree"
	PODMAN_INSTALLER_INSTALL  = "Install"
	PODMAN_INSTALLER_CLOSE    = "Close"

	installerBundleID    = "com.apple.installer"
	installerPodmanTitle = "Install Podman"
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
	delay.Delay(delay.XLONG)
	fmt.Print("with title")
	pInstaller, err := autoApp.LoadAppByTypeAndTitle(installerBundleID, installerPodmanTitle)
	logging.InitLogrus("", "", "")
	pInstaller.Print("", false)
	fmt.Print("after first load")
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pInstaller.Click(PODMAN_INSTALLER_CONTINUE, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	fmt.Print("after first continue")
	delay.Delay(delay.SMALL)
	pInstaller, err = pInstaller.Reload()
	fmt.Print("after first reload")
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pInstaller.Click(PODMAN_INSTALLER_CONTINUE, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	fmt.Print("after second continue")
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
	pdInstallerSec, err := autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pdInstallerSec.SetValueOnFocus(userPassword)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	delay.Delay(delay.XLONG)
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
