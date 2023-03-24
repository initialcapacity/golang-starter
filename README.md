# Golang starter

[![Build results](https://github.com/initialcapacity/golang-starter/workflows/build/badge.svg)](https://github.com/initialcapacity/golang-starter/actions)
[![codecov](https://codecov.io/gh/initialcapacity/golang-starter/branch/main/graph/badge.svg)](https://codecov.io/gh/initialcapacity/golang-starter)

An [application continuum](https://www.appcontinuum.io/) style example using Golang
that includes a single web application with 2 background workers.

* Basic web application
* Data analyzer
* Data collector

The example showcases on the below technologies -

* Language golang [Golang](https://go.dev/)
* Web Framework [Gorilla/mux](https://github.com/gorilla/mux)
* Build tool [Golang](https://go.dev/)
* Testing tools [Testify](https://github.com/stretchr/testify)
* Production [FreshCloud](https://www.freshcloud.com/) on Google's Cloud Platform

## Getting Started

Install the following prerequisites.

* [Go 1.20](https://go.dev)
* [Pack](https://buildpacks.io)
* [Docker Desktop](https://www.docker.com/products/docker-desktop)
* [Postgresql](https://www.postgresql.org/)

## For local development

Create a user and database.

```sql
drop database if exists starter_development;
drop user starter;

create user starter with password 'starter';
create database starter_development;
grant all privileges on database starter_development to starter;
```

Run the tests.

```bash
go clean -testcache && go test ./.../
```

Run the apps locally.

```bash
go run cmd/basicwebapp/app.go
```

Build and run the apps locally.

```bash
go install cmd/basicwebapp/app.go
~/go/bin/app
```

## For deployment

Build with Pack.

```bash
pack build golang-starter --builder heroku/buildpacks:20
```

Run with docker compose.

```bash
docker-compose up
````

That's a wrap for now.
