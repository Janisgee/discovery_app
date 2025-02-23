#!/bin/bash

DIR="$(dirname "${BASH_SOURCE[0]}")"

docker run --rm -v $DIR/../discovery-api:/discovery-api -w //discovery-api sqlc/sqlc generate
