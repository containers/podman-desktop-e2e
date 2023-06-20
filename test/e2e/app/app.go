package app

import (
	"fmt"

	autoApp "github.com/adrianriobo/goax/pkg/goax/app"
	podmanExtension "github.com/adrianriobo/podman-desktop-e2e/test/e2e/app/podman-extension"
)

func Cleanup() error {
	if err := cleanup(); err != nil {
		return err
	}
	return podmanExtension.Cleanup()
}

func Open(execPath string) error {
	return autoApp.Open(execPath)
}

func Close() error {
	pdApp, err := autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error closing the app: %v", err)
	}
	return pdApp.Click(APP_CLOSE, "", false)
}
