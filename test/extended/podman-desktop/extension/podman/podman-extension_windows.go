package podman

import (
	"github.com/adrianriobo/goax/pkg/os/windows/powershell"
	"github.com/adrianriobo/goax/pkg/util/delay"
	"github.com/adrianriobo/podman-desktop-e2e/test/extended/podman-desktop/util/ax"
)

const (
	installerInstall = "Install"
	installerClose   = "Close"
)

func cleanupSystem() error {
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

func runInstaller(userPassword string) error {
	delay.Delay(delay.MEDIUM)
	i, err := ax.GetForefront()
	if err != nil {
		return installerError(err)
	}
	if err := i.Click(installerInstall, delay.LONG); err != nil {
		return installerError(err)
	}
	if err := i.Click(installerClose, delay.SMALL); err != nil {
		return installerError(err)
	}
	return nil
}
