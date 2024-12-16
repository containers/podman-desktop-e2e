# podman-desktop-e2e

This project define a set of e2e tests around [podman desktop](https://github.com/podman-desktop/podman-desktop), this is a complementary set of e2e tests for those user stories which requires interaction beyond the podman desktop app itself with some third party apps (i.e installers from extensions).

## Overview

This project is intended for running without any extra dependecy, tests are self contained into a binary wich then will be run on a target host where podman dektop should be accessible.  

To accomplish this dependent-less runtime the project uses [goax](https://github.com/adrianriobo/goax) which uses OS native accessibility APIs to interact with UX elements (not only the ones from podman desktop but any other UX element rendered by the OS).

## Build

Currently the main two target OSs for the projects are Windows and MacOS: windows binary can be built on any platform, for building MacOS binary we need to build the binary on a MacOS machine (to ensure compatibilty it is recommended to build it on MacOS 12 Monterrey).  

Following commands will build the binary (`windows arm64 not supported`):  

```bash
# Build for mac amd64 (This need to be run on a MacOS)
ARCH=amd64 make build-darwin
# Binary will be located at
out/darwin-amd64/pd-e2e

# Build for mac arm64 (This need to be run on a MacOS)
ARCH=arm64 make build-darwin
# Binary will be located at
out/darwin-arm64/pd-e2e

# Build for windows amd64 
ARCH=amd64 make build-windows
# Binary will be located at
out/windows-amd64/pd-e2e.exe

```

## Run

The binary can be executed localy on the target hosts, it requires some parameters (on Windows):

* pdPath: Set the path where the podman desktop executable is located.
* pdUrl: Set the url where podman desktop executable will be downloaded.
* user-password: Set the password for the currrent user (this is needed for running installers which require elevated permissions).
* junit-filename: This is an optional parameter in case we want to set the name of the junit resulting file (Default: junit_report.xml).

Following command will run the tests on a windows host:  

```bash
pd-e2e.exe --pdPath /Users/rhqp/pd.exe --user-password MyPassword --junit-filename pd-e2e.xml 
```

Also the project is intended to be executed from a CI/CD system, [here](docs/running.md) is a full explanation on how to use it.

## Extend

The project is intended as an isolated project where e2e scenario can be defined and they would be implemented by interacting within the podman
desktop application through goax.

[Here](docs/extend.md) are some guidelines on how current suite of tests are defined and how this can be extended to cover more functionality from podman desktop.
