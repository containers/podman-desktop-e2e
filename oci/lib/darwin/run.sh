#!/bin/bash

# Validate required envs as params
mandatory_params () {
    local validate=1

    [[ -z "${PD_URL+x}" || -z "${PD_PATH+x}" ]] \
        && echo "PD_URL or PD_PATH are required" \
        && validate=0

    [[ -z "${USER_PASSWORD+x}" ]] \
        && echo "USER_PASSWORD required" \
        && validate=0

    [[ -z "${TARGET_FOLDER+x}" ]] \
        && echo "TARGET_FOLDER required" \
        && validate=0

    [[ -z "${RESULTS_FOLDER+x}" ]] \
        && echo "RESULTS_FOLDER required" \
        && validate=0

    [[ -z "${JUNIT_RESULTS_FILENAME+x}" ]] \
        && echo "JUNIT_RESULTS_FILENAME required" \
        && validate=0

    return $validate
}

install_pd () {
    pushd ${TARGET_FOLDER}
    curl -kL "${PD_URL}" -o pd.zip 
    unzip pd.zip
    popd 
}
 
if [[ ! mandatory_params ]]; then
    exit 1
fi

if [ "${DEBUG:-}" = "true" ]; then
  set -xuo 
fi

# Ensure no previous pd is running
pdpid=$(launchctl list | grep application.io.podmandesktop.PodmanDesktop | awk '{print $1}')
sudo kill -9 $pdpid

# Create results folder
mkdir -p "${HOME}/${TARGET_FOLDER}/${RESULTS_FOLDER}"

# Ensure we can execute pd-e2e
chmod +x $HOME/${TARGET_FOLDER}/pd-e2e

if [[ ! -z "${PD_URL+x}" ]]; then
    install_pd
    $HOME/${TARGET_FOLDER}/pd-e2e --user-password ${USER_PASSWORD} --junit-filename "${TARGET_FOLDER}/${RESULTS_FOLDER}/${JUNIT_RESULTS_FILENAME}" --pd-path "$HOME/${TARGET_FOLDER}/Podman Desktop.app" --screenshotspath "${RESULTS_FOLDER}"
else 
    $HOME/${TARGET_FOLDER}/pd-e2e --user-password ${USER_PASSWORD} --junit-filename "${TARGET_FOLDER}/${RESULTS_FOLDER}/${JUNIT_RESULTS_FILENAME}" --pd-path "${PD_PATH}" --screenshotspath "${RESULTS_FOLDER}"
fi

# Kill PD testing instance
pdpid=$(launchctl list | grep application.io.podmandesktop.PodmanDesktop | awk '{print $1}')
sudo kill -9 $pdpid
