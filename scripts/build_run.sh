#!/bin/bash

DIR="$(dirname "${BASH_SOURCE[0]}")"

$DIR/build.sh && $DIR/../discovery-api/dist/discovery_api.exe -env "$DIR/../discovery-api/.env"
