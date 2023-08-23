package podman

import (
	"fmt"

	"github.com/adrianriobo/goax/pkg/util/delay"
	"github.com/containers/podman-desktop-e2e/test/extended/podman-desktop/util/ax"
)

const (
	pdInstall  = "Install"
	pdPopUpYes = "Yes"
)

// Cleanup a previous installation for podmanextension
func CleanupSystem() error {
	return cleanupSystem()
}

// Install is accessible from dahsboard when podman is not isntalled on
// the system
func InstallPodman(pdApp *ax.AXApp, userPassword string) error {
	delay.Delay(delay.MEDIUM)
	if err := pdApp.Click(pdInstall, delay.MEDIUM); err != nil {
		return installerError(err)
	}
	if err := pdApp.Click(pdPopUpYes, delay.SMALL); err != nil {
		return installerError(err)
	}
	return runInstaller(userPassword)
}

// error format for installer error
func installerError(err error) error {
	return fmt.Errorf("error installing Podman: %v", err)
}
