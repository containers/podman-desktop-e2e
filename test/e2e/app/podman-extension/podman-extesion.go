package podmanextension

import (
	"fmt"

	autoApp "github.com/adrianriobo/goax/pkg/goax/app"
	"github.com/adrianriobo/goax/pkg/util/delay"
)

// Cleanup a previous installation for podmanextension
func Cleanup() error {
	return cleanup()
}

func Install() error {
	delay.Delay(delay.MEDIUM)
	pdApp, err := autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pdApp.Click(PODMAN_INSTALL, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	delay.Delay(delay.MEDIUM)
	pdApp, err = autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	err = pdApp.Click(PODMAN_INSTALL_POPUP_YES, "", true)
	if err != nil {
		return fmt.Errorf("error installing Podman: %v", err)
	}
	return nil
}

func Installer(userPassword string) error {
	return installer(userPassword)
}
