package context

type TestContextType struct {
	UserPassword          string
	JunitReportFilename   string
	AppPath               string
	ScreenshotsOutputPath string
}

var TestContext TestContextType

func SaveScreenshots() bool {
	return len(TestContext.ScreenshotsOutputPath) > 0
}
