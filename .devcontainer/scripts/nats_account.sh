#!/bin/bash

DIR="$( dirname -- "${BASH_SOURCE[0]}"; )";
echo "DIR IS $DIR"
DEVCONTAINER_DIR="$DIR/.."

sudo chown -Rh $USER:$USER $DEVCONTAINER_DIR/nsc

echo "Dumping NATS user creds file"
nsc --data-dir=$DEVCONTAINER_DIR/nsc/nats/nsc/stores generate creds -a IPAM -n USER > /tmp/user.creds

echo "Dumping NATS sys creds file"
nsc --data-dir=$DEVCONTAINER_DIR/nsc/nats/nsc/stores generate creds -a SYS -n sys > /tmp/sys.creds
