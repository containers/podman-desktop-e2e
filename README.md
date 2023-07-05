# podman-desktop-e2e

PoC for podman desktop e2e  

## Running remotely

podman-dektop-e2e binary is wrapped on a [container](https://github.com/adrianriobo/deliverest) which is responsible for copying it to the target host,
running the tests and got back the results.

Following snippet shows how this can be used, this sample was intended to be run from a folder holding the files with information for the target host:

```bash
# Darwin sample run
PD_E2E_V=1.1.0
podman run --rm -d --name pd-e2e-darwin \
    -e TARGET_HOST=$(cat host) \
    -e TARGET_HOST_USERNAME=$(cat username) \
    -e TARGET_HOST_KEY_PATH=/data/id_rsa \
    -e TARGET_FOLDER=pd-e2e \
    -e TARGET_RESULTS=pd-e2e-results.xml \
    -e TARGET_CLEANUP=true \
    -e OUTPUT_FOLDER=/data \
    -e DEBUG=true \
    -v $PWD:/data:z \
    quay.io/rhqp/podman-desktop-e2e:v${PD_E2E_V}-darwin-amd64 \
        USER_PASSWORD="$(cat userpassword)" \
        TARGET_FOLDER=pd-e2e \
        DEBUG=true \
        PD_PATH="/Users/$(cat username)/PodmanDesktop" \
        JUNIT_RESULTS_FILENAME=pd-e2e-results.xml \
        pd-e2e/run.sh

# Execution logs
podman logs -f pd-e2e-darwin

# Check results
cat pd-e2e-results.xml

# Windows sample run
PD_E2E_V=1.1.0
podman run --rm -d --name pd-e2e-windows \
    -e TARGET_HOST=$(cat host) \
    -e TARGET_HOST_USERNAME=$(cat username) \
    -e TARGET_HOST_KEY_PATH=/data/id_rsa \
    -e TARGET_FOLDER=pd-e2e \
    -e TARGET_RESULTS=pd-e2e-results.xml \
    -e OUTPUT_FOLDER=/data \
    -e DEBUG=true \
    -v $PWD:/data:z \
    quay.io/rhqp/podman-desktop-e2e:v${PD_E2E_V}-windows-amd64  \
        pd-e2e/run.ps1 \
            -targetFolder pd-e2e \
            -pdPath /Users/crcqe \
            -junitResultsFilename pd-e2e-results.xml 

# Execution logs
podman logs -f pd-e2e-windows

# Check results
cat pd-e2e-results.xml
```
