#!/bin/bash

export DB_CONNECT_MODE="DSN"

docker-compose start

GENERATE_DIR="./cmd/generate"

cd $GENERATE_DIR || exit

echo "Start Generating"
go run .
