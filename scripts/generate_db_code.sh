#!/bin/bash

DIR="$(dirname "${BASH_SOURCE[0]}")"

docker run --rm -v $DIR/../src:/src -w //src sqlc/sqlc generate
