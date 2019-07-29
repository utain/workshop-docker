#!/usr/bin/env bash
set -eo pipefail

case $1 in
  sit)
    # The '| cat' is to trick Node that this is an non-TTY terminal
    # then react-scripts won't clear the console.
    docker-compose -f ./build-sit.yml up --build| cat
    ;;
  dev)
    docker-compose -f ./start-dev.yml up --build| cat
    ;;
  prod)
    docker-compose -f ./docker-compose.yml up --build| cat
    ;;
  *)
    exec "$@"
    ;;
esac