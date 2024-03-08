package ax

import (
	"fmt"
	"time"

	autoApp "github.com/adrianriobo/goax/pkg/goax/app"
	"github.com/adrianriobo/goax/pkg/util/delay"
	"github.com/adrianriobo/goax/pkg/util/logging"
	"github.com/adrianriobo/goax/pkg/util/screenshot"
	"github.com/containers/podman-desktop-e2e/test/context"
)

// Represents any ax app
type AXApp struct{ ref *autoApp.App }

func GetForefront() (*AXApp, error) {
	a, err := autoApp.LoadForefrontApp()
	if err != nil {
		return nil, fmt.Errorf("error getting forefront app: %v", err)
	}
	return &AXApp{ref: a}, nil
}

func GetAppByTypeAndTitle(appType, appTitle string) (*AXApp, error) {
	a, err := autoApp.LoadAppByTypeAndTitle(appType, appTitle)
	if err != nil {
		return nil, fmt.Errorf("error getting the app %s with title %s: %v", appType, appTitle, err)
	}
	return &AXApp{ref: a}, nil
}

// Click on a clickable element and wait an amount of delay
func (a *AXApp) Click(element string, delayAmount time.Duration) error {
	return a.ClickWithType(element, "", delayAmount)
}

// Click on a clickable element by type and wait an amount of delay
func (a *AXApp) ClickWithType(element, elementType string, delayAmount time.Duration) error {
	r, err := a.ref.Reload()
	if err != nil {
		return fmt.Errorf("error reloading the application: %v", err)
	}
	err = r.Click(element, elementType, true)
	if err != nil {
		return fmt.Errorf("error clicking on %s: %v", element, err)
	}
	delay.Delay(delayAmount)
	if context.SaveScreenshots() {
		if err := screenshot.CaptureScreen(context.TestContext.ScreenshotsOutputPath,
			fmt.Sprintf("click-%s%s", element, elementType)); err != nil {
			logging.Errorf("error capturing the screenshot: %v", err)
		}
	}
	return nil
}

// Click on a clickable element by type and wait an amount of delay
func (a *AXApp) ExistsWithType(element, elementType string) (bool, error) {
	r, err := a.ref.Reload()
	if err != nil {
		return false, fmt.Errorf("error reloading the application: %v", err)
	}
	if context.SaveScreenshots() {
		if err := screenshot.CaptureScreen(context.TestContext.ScreenshotsOutputPath,
			fmt.Sprintf("exists-%s%s", element, elementType)); err != nil {
			logging.Errorf("error capturing the screenshot: %v", err)
		}
	}
	return r.Exists(element, elementType, false)
}

func (a *AXApp) SetValueOnFocus(value string, delayAmount time.Duration) error {
	if err := a.ref.SetValueOnFocus(value); err != nil {
		return err
	}
	delay.Delay(delayAmount)
	return nil
}

func (a *AXApp) Print() {
	a.ref.Print("", false)
}
