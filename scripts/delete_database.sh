#!/bin/bash

# Stop the PostgreSQL container
docker stop discovery_db

# Remove the PostgreSQL container
docker rm discovery_db

# Remove the associated volume
docker volume rm discovery_app_pg_data
