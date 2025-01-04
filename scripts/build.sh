#!/bin/bash

DIR="$(dirname "${BASH_SOURCE[0]}")"

cd $DIR/../src
go get
go build -o ../dist/discovery_app.exe
