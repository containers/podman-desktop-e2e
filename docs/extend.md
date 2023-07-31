# Testing framework

Testing framework and E2E tests are written in [Go](https://go.dev/) using [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) frameworks to cover Podman Desktop functionality.

It also uses [Goax](https://github.com/adrianriobo/goax) as an abstration library to interact with Podman Desktop and any other UX element using OS native accessibility API.

## Testing structure

Podman test definition can be found at [e2e test folder](./../test/e2e/e2e_podman/). From there it is expected any podman-desktop component will have its own file and will define all the functionality to be tested.  

As an example for [podman-extension](./../test/e2e/e2e_podman/podman-extension_test.go) we describe the funcionality to install podman through its podman desktop extension:

```golang
var _ = ginkgo.Describe("podman-extension [extension-podman]", func() {
    ginkgo.It("can be installed", func() {
        err := podmanExtension.InstallPodman(PDHandler.AXApp, TestContext.UserPassword)
        gomega.Expect(err).NotTo(gomega.HaveOccurred())
    })
})

```

Previous snipped showcase how the actions available within Podman Desktop are executed by an application handler. The application handler is implemented at [podman desktop handler](./../test/extended/podman-desktop/) and it is expected each extension functionality will be also included as an specific extension at [extensions](./../test/extended/podman-desktop/extension/) folder.

### Goax API

Goax framework is used to interact with application itself. The framework works by getting an application from the X session of the Operating System and get all refereces to accessible elements contained in it.  

For the application handler the application is [open](./../test/extended/podman-desktop/podman-desktop.go#L30) and then it is [loaded]((./../test/extended/podman-desktop/podman-desktop.go#L35)) as an [accessible application](https://github.com/adrianriobo/goax/blob/main/pkg/goax/axapp/axapp.go) with that reference we can interact which any accessible element on it. Al possible functions we can invoke on the application then:

* Reload: If an element is clicked on the application and it change the elements on it, we need to reload it to get the new references for the new visual elements.  
  
* Exists: Check if an elemet exists on the application by its name and type.  

* Click: Allows to click on an element from the application based in its name and type.

* ClickWithOrder: Same as Click, but in case multiple elements we can define which of them by its order.

* SetValue: Allows to set a value on an element (i.e fill a textbox) from the application based in its name and type.

* SetValueWithOrder: Same as SetValue, but in case multiple elements we can define which of them by its order.

* SetValueOnFocus: If some element on the application has the focus and is an editable element, this funcion allows to set a value on it. (Only supported for darwin platform)
