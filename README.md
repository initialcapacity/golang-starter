# Golang starter

[![Build results](https://github.com/initialcapacity/golang-starter/workflows/build/badge.svg)](https://github.com/initialcapacity/golang-starter/actions)
[![codecov](https://codecov.io/gh/initialcapacity/golang-starter/branch/main/graph/badge.svg)](https://codecov.io/gh/initialcapacity/golang-starter)

An [application continuum](https://www.appcontinuum.io/) style example using Golang
that includes a single web application with 2 background workers.
Deployed via [Fresh Cloud](https://www.freshcloud.com/).

* Basic web application
* Data analyzer
* Data collector

The example showcases on the below technologies -

* Language golang [Golang](https://go.dev/)
* Web Framework [Gorilla/mux](https://github.com/gorilla/mux)
* Build tool [Golang](https://go.dev/)
* Testing tools [Testify](https://github.com/stretchr/testify)
* Production [FreshCloud](https://www.freshcloud.com/) on Google's Cloud Platform

```bash
go get github.com/gorilla/mux
go get github.com/stretchr/testify
```

## Getting Started

Install the following prerequisites.

* [Go 1.18](https://go.dev)
* [Pack](https://buildpacks.io)
* [Docker Desktop](https://www.docker.com/products/docker-desktop)
* [Postgresql](https://www.postgresql.org/)

Build with Pack.

```bash
pack build golang-starter --builder heroku/buildpacks:20
```

Run with docker compose.

```bash
docker-compose up
````

## Development

Ensure postgres works locally.

```bash
"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
```

Run the tests.

```bash
go clean -testcache && go test ./.../
```

## Deployment

Fresh cloud deployment and pipeline files are located in `deployments`.

That's a wrap for now.
