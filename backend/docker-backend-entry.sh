#!/bin/sh

# Replaces environment variables in index.html by their value

CONFIG_DIR="/go/src/app/"

envsubst < "$CONFIG_DIR/config.json.template" | tee "$CONFIG_DIR/config.json" > /dev/null

echo "Starting backend"

[ -z "$@" ] && rezept "$CONFIG_DIR/config.json" || $@