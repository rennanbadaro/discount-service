#! /bin/bash

# export env vars from .env
export $(cat .env | xargs)

golang-migrate \
    -path db/migrations \
    -database postgres://$DB_USERNAME:$DB_PASSWORD@127.0.0.1:$DB_PORT/hasher_store\?sslmode=disable \
    up
