# Running model

Project is intended to be run within a CI/CD system, as so the binary is distributed within a container to easly integrate within any CI/CD system or even to execute on any container runtime targetting the remote host where tests will be executed.

As a side note it is important to know that in order to access the X elements, an X session should exists on the target host. 

For mac systems it is enough if the user has an autologin property set but on windows this is not enough, and an active X session is required. To solve this issue when running the project remotely on windows it is advisable to run a secondary container to create a [fake rdp connection](https://github.com/adrianriobo/frdp).

## OCI container

podman-dektop-e2e binary is wrapped on a [container](https://github.com/adrianriobo/deliverest) which is responsible for copying it to the target host, running the tests and got back the results.

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

## Tekton task

This project includes a [task definition](./../tkn/task.yaml) to include its execution as part of Openshift Pipeline. The task includes all the required parameters to run the container connecting to the remote target host where podman-desktop will be tested.  

The task is publised on quay at https://quay.io/repository/rhqp/podman-desktop-e2e-tkn and its definition can be used directly using the bundle resolver:

```yaml
...
  tasks:
  - name: podman-desktop-e2e
    taskRef:
      resolver: bundles 
      params:
      - name: bundle
        value: quay.io/rhqp/podman-desktop-e2e-tkn:v0.1
      - name: name
        value: podman-desktop-e2e
      - name: kind
        value: task
    params:
    - name: os
...
```

Within the task we include a [sidecar running a fakerdp](./../tkn/task.yaml#L118) connection to emulate a X session on Windows machines.
