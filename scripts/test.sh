#!/usr/bin/env bash

###
# Runs the service tests.
###

function read_env_file {
  set -a
  [ -f .env.test ] && source .env.test
  set +a
}

##
# Main function
##
function main {
  local db_container_id

  # Read test env file
  read_env_file

  # Runs postgres Docker image in the background.
  db_container_id=$(docker run \
  --name mimir_db_test \
  -e POSTGRES_DB="${DB_NAME}" \
  -e POSTGRES_PASSWORD="${DB_PASSWORD}" \
  -e POSTGRES_USER="${DB_USER}" \
  -d \
  -p "${DB_PORT}":5432 \
  postgres:12-alpine)
  echo "Running PostgreSQL container: ${db_container_id}"

  # Wait for container to start.
  sleep 2

  # Run tests.
  go test -v ./...

  # Force remove database container.
  echo "Removing PostgreSQL container: ${db_container_id}"
  docker rm -f "${db_container_id}"
}

# And so, it begins...
main
