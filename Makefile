
PODMAN_DESKTOP_VERSION ?= 1.1.0
CONTAINER_MANAGER ?= podman

# Image URL to use all building/pushing image targets
IMG ?= quay.io/rhqp/podman-desktop-e2e:v${PODMAN_DESKTOP_VERSION}
E2E_BINARY ?= pd-e2e

# Tekton task
TKN_TASK_VERSION ?= 0.1
TKN_IMG ?= quay.io/rhqp/podman-desktop-e2e-tkn:v${TKN_TASK_VERSION}


BUILD_DIR ?= out
NATIVE_GOARCH := $(shell env -u GOARCH go env GOARCH)
ARCH ?= $(NATIVE_GOARCH)

# https://golang.org/cmd/link/
LDFLAGS := $(VERSION_VARIABLES) -extldflags='-static' ${GO_EXTRA_LDFLAGS}

TOOLS_DIR := tools
include tools/tools.mk

.PHONY: clean 
clean: 
	rm -rf $(BUILD_DIR)

.PHONY: build
build:
	go test -v test/e2e/e2e_podman/suite_test.go test/e2e/e2e_podman/podman-extension_test.go -c -o $(BUILD_DIR)/linux-amd64/pd-e2e

.PHONY: cross
cross: clean $(BUILD_DIR)/windows-amd64/pd-e2e.exe

.PHONY: build-windows
build-windows: clean $(BUILD_DIR)/windows-${ARCH}/pd-e2e.exe

$(BUILD_DIR)/windows-${ARCH}/pd-e2e.exe: $(SOURCES)
	CC=clang GOARCH=${ARCH} GOOS=windows go test -v test/e2e/e2e_podman/suite_test.go test/e2e/e2e_podman/podman-extension_test.go \
		-c -o $(BUILD_DIR)/windows-${ARCH}/pd-e2e.exe 

.PHONY: build-darwin
build-darwin: clean $(BUILD_DIR)/darwin-${ARCH}/pd-e2e

$(BUILD_DIR)/darwin-${ARCH}/pd-e2e:
	CGO_ENABLED=1 CC=clang GOARCH=${ARCH} GOOS=darwin go test -v test/e2e/e2e_podman/suite_test.go test/e2e/e2e_podman/podman-extension_test.go \
		-c -o $(BUILD_DIR)/darwin-${ARCH}/pd-e2e 
    
.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

# Build the container image
.PHONY: oci-build
oci-build: 
	${CONTAINER_MANAGER} build -t ${IMG}-${OS}-${ARCH} -f oci/Containerfile --build-arg=OS=${OS} --build-arg=ARCH=${ARCH} --build-arg=E2E_BINARY=${E2E_BINARY} oci

# Build the container image
.PHONY: oci-push
oci-push: 
	${CONTAINER_MANAGER} push ${IMG}-${OS}-${ARCH}

# Create tekton task bundle
.PHONY: tkn-push
tkn-push: install-out-of-tree-tools
	$(TOOLS_BINDIR)/tkn bundle push $(TKN_IMG) -f tkn/task.yaml