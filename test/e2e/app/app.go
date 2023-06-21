package app

import (
	"fmt"

	autoApp "github.com/adrianriobo/goax/pkg/goax/app"
	"github.com/adrianriobo/goax/pkg/util/delay"
	"github.com/adrianriobo/goax/pkg/util/logging"
	podmanExtension "github.com/adrianriobo/podman-desktop-e2e/test/e2e/app/podman-extension"
)

func Cleanup() error {
	if err := cleanup(); err != nil {
		logging.Errorf("error cleaning up system from podman: %v", err)
		return err
	}
	if err := podmanExtension.Cleanup(); err != nil {
		logging.Errorf("error cleaning up system from podman extension: %v", err)
		return err
	}
	return nil
}

func Open(execPath string) error {
	err := autoApp.Open(execPath)
	// We open remotely so we wait for a bit
	delay.Delay(delay.LONG)
	return err
}

func Close() error {
	pdApp, err := autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error closing the app: %v", err)
	}
	return pdApp.Click(APP_CLOSE, "", false)
}
