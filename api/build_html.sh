#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0) && pwd)

for f in $SCRIPT_DIR/functions/*.sh; do source "$f"; done
deps_check 'redoc-cli' 'required redoc-cli, npm i -g redoc-cli'

if [ ! -e $SCRIPT_DIR/openapi.yaml ]; then
    echo "not found specification file."
    exit 1
fi
deps_check 'redoc-cli' 'required redoc-cli, npm i -g redoc-cli'
redoc-cli bundle $SCRIPT_DIR/openapi.yaml --options.showExtensions --options.disableSearch --title "Data API" -o $SCRIPT_DIR/openapi.html

exit 0