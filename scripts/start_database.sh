#!/bin/bash

docker run \
  --name discovery_db \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_USER=posgres \
  -e POSTGRES_DB=discovery_app \
  -p 5433:5432 \
  -v discovery_app_pg_data:/var/lib/postgresql/data \
  -d \
  postgres
