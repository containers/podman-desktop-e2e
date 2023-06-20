package podman

import (
	podmanExtension "github.com/adrianriobo/podman-desktop-e2e/test/e2e/app/podman-extension"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("podman-extension [extension-podman]", func() {
	ginkgo.It("can be installed", func() {
		err := podmanExtension.Install()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		// On darwin we need the user password to run the installation
		err = podmanExtension.Installer(TestContext.UserPassword)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.It("some value can be configured", func() {

	})
})
