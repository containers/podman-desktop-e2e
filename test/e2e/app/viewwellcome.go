package app

import (
	"fmt"

	autoApp "github.com/adrianriobo/goax/pkg/goax/app"
)

// On Wellcome page we should have telemetry as enable by default
// we change to disable and go to podman
func DisableTelemetryOnWellcomePage() error {
	pdApp, err := autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error handling the Welcome Page: %v", err)
	}
	exists, err := pdApp.Exists(WELLCOME_PAGE_TELEMETRY_ENABLE, "checkbox", false)
	if err != nil {
		return fmt.Errorf("error handling the Welcome Page: %v", err)
	}
	if !exists {
		return fmt.Errorf("expect to have a check for %s, but we could not find it", WELLCOME_PAGE_TELEMETRY_ENABLE)
	}
	return pdApp.Click(WELLCOME_PAGE_TELEMETRY_ENABLE, "checkbox", false)
}

func GoToPodman() error {
	pdApp, err := autoApp.LoadForefrontApp()
	if err != nil {
		return fmt.Errorf("error going to podman from Welcome Page: %v", err)
	}
	return pdApp.Click(WELLCOME_PAGE_GO_TO_PD, "", false)
}
