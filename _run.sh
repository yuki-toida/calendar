#!/bin/sh

# node process kill
killall node

# docker run
docker-compose up -d

cd assets
yarn install
yarn watch &
cd ../

# dep ensure -update -v
dep ensure

export ENV=local
export GOOGLE_APPLICATION_CREDENTIALS="./cred/gcs.json"
go run main.go
