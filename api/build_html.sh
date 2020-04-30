#!/bin/bash
SCRIPT_DIR="$(cd $(dirname $0) && pwd)"
cd $SCRIPT_DIR

redoc-cli bundle ./openapi.yaml --options.disableSearch --title "Data API" -o openapi.html