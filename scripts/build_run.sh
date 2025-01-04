#!/bin/bash

DIR="$(dirname "${BASH_SOURCE[0]}")"

$DIR/build.sh && $DIR/../dist/discovery_app.exe
