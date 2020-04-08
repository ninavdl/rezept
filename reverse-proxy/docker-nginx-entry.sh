#!/bin/sh

# Replaces environment variables in nginx.conf by their value

envsubst < "default.conf.template" | tee "/etc/nginx/conf.d/default.conf" > /dev/null

echo "Starting nginx"

[ -z "$@" ] && nginx -g 'daemon off;' || $@