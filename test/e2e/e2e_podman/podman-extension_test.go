package podman

import (
	"github.com/containers/podman-desktop-e2e/test/context"
	podmanExtension "github.com/containers/podman-desktop-e2e/test/extended/podman-desktop/extension/podman"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("podman-extension [extension-podman]", func() {
	ginkgo.It("can be installed", func() {
		err := podmanExtension.InstallPodman(PDHandler.AXApp, context.TestContext.UserPassword)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.It("some value can be configured", func() {

	})
})
