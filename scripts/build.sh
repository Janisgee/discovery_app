#!/bin/bash

DIR="$(dirname "${BASH_SOURCE[0]}")"

cd $DIR/../discovery-api
go get
go build -o ./dist/discovery_api.exe
