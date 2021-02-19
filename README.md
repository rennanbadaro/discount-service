## Discount Service

Calculates proper discount for certain product and user.

## Requirements

- Go v1.15
- [https://github.com/golang-migrate/migrate](golang-migrate) v4.14
- Docker
- Docker Compose

## Usage

Before running the running the project make sure you've met the requirements described above. Regarding the `golang-migrate` tool, after downloading it, add it to your $PATH with the alias `golang-migrate`. This will be needed to setup up your local DB.

To run the project locally run:

```sh
cp .env.example && chmod +x ./scripts/**
docker-compose up
```

This will spin up a API and a Postgres container. On another terminal, setup the DB for you local usage with the script below. It will run migrations to create the needed tables and seed them with some fake data

```sh
./scripts/setup-dev-db.sh
```

In case you want to undo this setup, simply run:

```sh
./scripts/teardown-dev-db.sh
```

## Tests

```sh
go test ./discount
```

## Proto files

The project already contains generated code. The source for the proto files lives in the repo [proto-graal](https://github.com/rennanbadaro/proto-graal). In case there's an update in `proto-graal` that should have an impact on this project, the code can be regenerated any time by running the shell script `./scripts/build-protos.sh`. It will clone the latest version of `proto-graal` and place the code into the proper directory.

## Development Notes

The `discount-service` takes a simpler path than its counterpart [product-service](https://github.com/rennanbadaro/product-service) in terms of architecture. This was actually intented as an attempt to build a more idiomatic Go application, but specially because of its few and straightforward components.

The design tries to express a separation of 3 layers: application, domain, infrastructure which are directly related to the directories/packages `cmd`, `discount` and `infrastructure`, respectively.

A [Trello Board](https://trello.com/b/33FKh8uj/hasher-store) was used to keep track of all the tasks.
