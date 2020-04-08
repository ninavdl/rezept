#!/bin/sh

# Replaces environment variables in index.html and nginx default.conf by their value

envsubst < "/usr/share/nginx/html/index.html.template" | tee "/usr/share/nginx/html/index.html" > /dev/null

echo "Starting nginx"

[ -z "$@" ] && nginx -g 'daemon off;' || $@