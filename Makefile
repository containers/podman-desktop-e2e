
BUILD_DIR ?= out
ARCH ?= amd64

# https://golang.org/cmd/link/
LDFLAGS := $(VERSION_VARIABLES) -extldflags='-static' ${GO_EXTRA_LDFLAGS}

.PHONY: clean 
clean: 
	rm -rf $(BUILD_DIR)

.PHONY: build
build:
	go test -v test/e2e/e2e_podman/suite_test.go test/e2e/e2e_podman/podman-extension_test.go -c -o $(BUILD_DIR)/linux-amd64/pd-e2e

.PHONY: cross
cross: clean $(BUILD_DIR)/windows-amd64/pd-e2e.exe

.PHONY: build-windows
build-windows: clean $(BUILD_DIR)/windows-amd64/pd-e2e.exe

$(BUILD_DIR)/windows-amd64/pd-e2e.exe: $(SOURCES)
	CC=clang GOARCH=amd64 GOOS=windows go test -v test/e2e/e2e_podman/suite_test.go test/e2e/e2e_podman/podman-extension_test.go \
		-c -o $(BUILD_DIR)/windows-amd64/pd-e2e.exe 

.PHONY: build-darwin
build-darwin: clean $(BUILD_DIR)/darwin-${ARCH}/pd-e2e

$(BUILD_DIR)/darwin-${ARCH}/pd-e2e:
	CGO_ENABLED=1 CC=clang GOARCH=${ARCH} GOOS=darwin go test -v test/e2e/e2e_podman/suite_test.go test/e2e/e2e_podman/podman-extension_test.go \
		-c -o $(BUILD_DIR)/darwin-${ARCH}/pd-e2e 
    
.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor