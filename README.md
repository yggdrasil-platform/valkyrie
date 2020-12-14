# Valkyrie

Provides information about the other services.

#### Table of contents

* [Introduction](#introduction)
    * [Project structure](#project-structure)
* [Development](#development)
    * [1. Prerequisites](#1-prerequisites)
    * [2. Setup](#2-setup)
    * [3. Running locally](#3-running-locally)
    * [4. Regenerating GraphQL schema](#4-regenerating-graphql-schema)
* [Testing](#testing)
    * [1. Running tests](#1-running-tests)
* [Deployment](#deployment)
    * [1. Building the image](#1-building-the-image)
    * [2. Running the image](#2-running-the-image)

## Introduction

### Project structure

Below is a quick outline of the structure of the app:

```text
.
├── cmd                        # Go apps.
|   ├── server
|   │   └── main.go            # Entry point for the app.
│   └── ...
├── data                       # A collection of seeds to add to the database.
|   ├── model.json
│   └── ...
├── pkg                        # A collection of reusable modules.
|   ├── application
|   |   └── application.go     # Sets up the configuration for the app.
|   ├── cleanup
|   |   └── cleanup.go         # Gracefully cleans up the app.
|   ├── config
|   |   └── config.go          # Struct that defines the application config.
|   ├── database
|   |   ├── database.go        # Sets up the database connection and runs migrations/seeds.
|   |   ├── migration.go       # Migrations helper.
|   |   └── seed.go            # Seed helper.
|   ├── error
|   |   └── code.go            # All the error codes and the descriptions.
|   ├── graphql
|   |   └── graphql_gen.go     # The auto-generated GraphQL helper.
|   ├── handler                # Handles HTTP requests and responses.
|   |   ├── somehandler.go
|   |   └── ...
|   ├── logger                 # Module to pretty print log messages.
|   |   └── ...
|   ├── middleware             # Custom middleware.
|   |   ├── somemiddleware.go
|   |   └── ...
|   ├── model                  # Contains the database and GraphQL models.
|   |   ├── somemodel.go
|   |   └── ...
|   ├── resolver               # Contains the GraphQL resolvers.
|   |   ├── resolver.go        # The main resolver struct.
|   |   ├── someresolver.go
|   |   └── ...
|   └── router
|   |   └── router.go          # Creates the router and sets up the handlers.
|   ├── server
|   |   ├── request.go         # Request structs.
|   |   ├── response.go        # Response structs.
|   |   └── server.go          # Starts the server.
|   ├── service                # Reusable code such as database operations
|   |   ├── someservice.go
|   |   └── ...
|   └── ...
├── schema                     # A collection of GraphQL schema files.
|   ├── someschema.graphqls
│   └── ...
└── ...
```

## Development

### 1. Prerequisites

* Install [Go 1.14+](https://golang.org/dl/).
* Install [Docker](https://docs.docker.com/get-docker/).

### 2. Setup

1. Install the dependencies:
```bash
go get
```

### 3. Running locally

1. Simply run:
```bash
./scripts/start.sh
```

2. You can check the API is running using the following cURL command:
```shell script
curl -X GET http://localhost:${PORT}/healthcheck
```

### 4. Regenerating GraphQL schema

This service uses the excellent [gqlgen](https://gqlgen.com). However, any new or edited schemas need to regenerate the GraphQL server code and create/update the resolvers and models.

To do this, simply run:
```shell
go generate ./...
```

## Testing

### 1. Running tests

1. Simply run the command:
```bash
./scripts/test.sh
```

## Deployment

### 1. Building the image

When building the Docker image, we want to inject env vars at build time, as the [`Dockerfile`](./Dockerfile) injects the build args as env vars into the container.
```bash
docker build \
-t kieranroneill/valkyrie \
--build-arg db_host=edda \
--build-arg db_name=valkyrie \
--build-arg db_password=password \
--build-arg db_port=5432 \
--build-arg db_user=admin \
--build-arg env=production \
--build-arg port=3000 \
--build-arg service_name=valkyrie \
--build-arg version=$(<VERSION) \
.
```

#### 2. Running the image

```bash
docker run \
--name valkyrie \
-it \
-p 1337:${PORT} \
kieranroneill/valkyrie:latest
```
