package podman

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/reporters"
	"github.com/onsi/ginkgo/v2/types"
	"github.com/onsi/gomega"
	"github.com/spf13/pflag"

	"github.com/containers/podman-desktop-e2e/test/context"
	podmanDesktop "github.com/containers/podman-desktop-e2e/test/extended/podman-desktop"
)

func TestMain(m *testing.M) {
	RegisterCommonFlags(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	AfterReadingAllFlags()
	os.Exit(m.Run())
}

func TestE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	suiteConfig, reporterConfig := CreateGinkgoConfig()
	ginkgo.RunSpecs(t, "E2EPodman suite", suiteConfig, reporterConfig)
}

func CreateGinkgoConfig() (types.SuiteConfig, types.ReporterConfig) {
	// fetch the current config
	suiteConfig, reporterConfig := ginkgo.GinkgoConfiguration()
	// Turn on EmitSpecProgress to get spec progress (especially on interrupt)
	suiteConfig.EmitSpecProgress = true
	// Randomize specs as well as suites
	suiteConfig.RandomizeAllSpecs = true
	// Turn on verbose by default to get spec names
	reporterConfig.Verbose = true
	// Disable skipped tests unless they are explicitly requested.
	if len(suiteConfig.FocusStrings) == 0 && len(suiteConfig.SkipStrings) == 0 {
		suiteConfig.SkipStrings = []string{`\[Flaky\]|\[Feature:.+\]`}
	}
	return suiteConfig, reporterConfig
}

func RegisterCommonFlags(flags *flag.FlagSet) {
	flags.StringVar(&context.TestContext.AppPath, "pd-path", "", "Set the user password to be used within the tests.")
	flags.StringVar(&context.TestContext.UserPassword, "user-password", "", "Set the user password to be used within the tests.")
	flags.StringVar(&context.TestContext.JunitReportFilename, "junit-filename", "", "Set the filename for the junit report.")
	flags.StringVar(&context.TestContext.ScreenshotsOutputPath, "screenshotspath", "", "Set the path to save screenshots.")
}

func AfterReadingAllFlags() {
	ginkgo.ReportAfterSuite("Podman Desktop e2e JUnit report", writeJUnitReport)
}

// INFO Inspired by https://github.com/kubernetes/kubernetes/blob/07315d10b3718973e5ebcc61cbf0fba8a6ec58a9/test/e2e/framework/test_context.go#LL535C1-L573C2
func writeJUnitReport(report ginkgo.Report) {
	filename := context.TestContext.JunitReportFilename
	if len(filename) == 0 {
		filename = generateDefaultJunitReportName()
	}
	if err := reporters.GenerateJUnitReportWithConfig(report,
		filename,
		reporters.JunitReportConfig{
			OmitFailureMessageAttr:    true,
			OmitCapturedStdOutErr:     true,
			OmitTimelinesForSpecState: types.SpecStatePassed,
		}); err != nil {
		fmt.Printf("error setting the junit reporter: %v", err)
	}
}

// TODO add timestamp
func generateDefaultJunitReportName() string {
	return "junit_report.xml"
}

var PDHandler *podmanDesktop.PDApp

var _ = ginkgo.BeforeSuite(func() {
	// Cleanup system ref to PodmanDesktop to ensure fresh env
	if err := podmanDesktop.CleanupSystem(); err != nil {
		fmt.Printf("error cleaning up the system %v", err)
	}
	// gomega.Expect(err).NotTo(gomega.HaveOccurred())
	// Open the app with param from exec
	var err error
	PDHandler, err = podmanDesktop.Open(context.TestContext.AppPath)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	// First run will show welcome page
	err = PDHandler.WelcomePageDisableTelemetry()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	// Go to Podman
	err = PDHandler.WelcomePageGoToPodman()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

})

var _ = ginkgo.AfterSuite(func() {
	// err := podmanDesktop.Close()
	// gomega.Expect(err).NotTo(gomega.HaveOccurred())
})
