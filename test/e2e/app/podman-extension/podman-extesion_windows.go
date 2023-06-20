package podmanextension

import (
	"fmt"
	autoApp "github.com/adrianriobo/goax/pkg/goax/app"
	"github.com/adrianriobo/goax/pkg/os/windows/powershell"
	"github.com/adrianriobo/goax/pkg/util/delay"
)

func cleanup() error {
	ps := powershell.New()
	// Start-Process powershell -verb runas -ArgumentList "Remove-Item 'C:\Program Files\RedHat' -Force"
	removePodmanHome := "Start-Process powershell -verb runas -WindowStyle Hidden -Wait -ArgumentList \"Remove-Item 'C:\\Program Files\\RedHat\\Podman' -Recurse -Force -Erroraction silentlycontinue \""
	// removePodmanHome := "Remove-Item \"$env:PROGRAMFILES\\RedHat\\Podman\" -Recurse -Force -Erroraction silentlycontinue"
	ps.Execute(removePodmanHome)
	// if err != nil {
	// 	return err
	// }
	removePodmanShare := "Remove-Item \"$env:USERPROFILE\\.local\\share\\containers\" -Recurse -Force -Erroraction silentlycontinue"
	ps.Execute(removePodmanShare)
	// if err != nil {
	// 	return err
	// }
	removePodmanConfig := "Remove-Item \"$env:USERPROFILE\\.config\\containers\" -Recurse -Force -Erroraction silentlycontinue"
	ps.Execute(removePodmanConfig)
	return nil
}

func installer(userPassword string) error {
	delay.Delay(delay.MEDIUM)
	pInstaller, err := autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pInstaller.Click(PODMAN_INSTALLER_INSTALL, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	delay.Delay(delay.LONG)
	pInstaller, err = pInstaller.Reload()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	return pInstaller.Click(PODMAN_INSTALLER_CLOSE, "", true)
}
