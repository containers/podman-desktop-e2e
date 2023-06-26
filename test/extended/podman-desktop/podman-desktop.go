package podmandesktop

import (
	"fmt"

	autoApp "github.com/adrianriobo/goax/pkg/goax/app"
	"github.com/adrianriobo/goax/pkg/util/delay"
	"github.com/adrianriobo/goax/pkg/util/logging"
	podmanExtension "github.com/adrianriobo/podman-desktop-e2e/test/extended/podman-desktop/extension/podman"
	"github.com/adrianriobo/podman-desktop-e2e/test/extended/podman-desktop/util/ax"
)

type PDApp struct {
	*ax.AXApp
}

func CleanupSystem() error {
	if err := cleanup(); err != nil {
		logging.Errorf("error cleaning up system from podman: %v", err)
		return err
	}
	if err := podmanExtension.CleanupSystem(); err != nil {
		logging.Errorf("error cleaning up system from podman extension: %v", err)
		return err
	}
	return nil
}

func Open(execPath string) (*PDApp, error) {
	if err := autoApp.Open(execPath); err != nil {
		return nil, fmt.Errorf("error opening the podman desktop executable at %s: %v", execPath, err)
	}
	// We open remotely so we wait for a bit
	delay.Delay(delay.LONG)
	a, err := ax.GetForefront()
	if err != nil {
		return nil, fmt.Errorf("error opening the podman desktop executable at %s: %v", execPath, err)
	}
	return &PDApp{a}, nil
}

// On Wellcome page we should have telemetry as enable by default
// we change to disable and go to podman
func (p *PDApp) WellcomePageDisableTelemetry() error {
	exists, err := p.ExistsWithType(wellcomePageEnableTelemetry, "checkbox")
	if err != nil || !exists {
		return fmt.Errorf("error disabling telemetry :%v", err)
	}
	if err := p.ClickWithType(wellcomePageEnableTelemetry, "checkbox", delay.SMALL); err != nil {
		return fmt.Errorf("error disabling telemetry: %v", err)
	}
	return nil
}

func (p *PDApp) WellcomePageGoToPodman() error {
	if err := p.Click(wellcomePageGoToPodmanDesktop, delay.SMALL); err != nil {
		return fmt.Errorf("error going to wellcome page: %v", err)
	}
	return nil
}

func (p *PDApp) Close() error {
	if err := p.Click(appClose, delay.SMALL); err != nil {
		return fmt.Errorf("error closing the app: %v", err)
	}
	return nil
}
