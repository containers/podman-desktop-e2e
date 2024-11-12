#!/bin/bash
podmanMachine="podman-machine"
remoteMachine="remote-machine"

echo "Preparing remote podman machine..."

podman machine init --now $podmanMachine

podman pull ghcr.io/podmandesktop-ci/alpine-remote

podman image exists podmandesktop-ci/alpine-remote && echo "Image exists... " || { echo "image does not exists"; exit 1 }

# Parse the podman system connection json to get default connection

name=$(podman system connection ls --format json | jq '.[] | select(.Default==true)' | jq -r '.Name')
uri=$(podman system connection ls --format json | jq '.[] | select(.Default==true)' | jq -r '.URI')
identity=$(podman system connection ls --format json | jq '.[] | select(.Default==true)' | jq -r '.Identity')

echo "Default connection - Name: $name, URI: $uri, Identity: $identity"

podman system connection rm $podmanMachine
podman system connection rm $podmanMachine-root

# remove default connections and other podman related file
if [[ "$(uname)" = *Linux* ]]; then 
    rm -rf /run/user/1000/podman/*podman-machine*
fi

# Do not remove ~/.local/share/containers/podman as there are keys to the machine

# remove other configuration files from USERPROFILE\.config\containers
rm -rf ~/.config/containers

# create a connection from previous information
podman system connection add $remoteMachine --identity $identity $uri

# check connection
podman system connection --format json

# run the remote e2e tests...
