# Golang starter example

An application continuum style example application using Golang that includes a single web application with 2 background
workers. Deployed via Fresh Cloud.

* Basic web application
* Data analyzer
* Data collector

The example showcases on the below technologies -

* Language golang
* 1.18 generics
* Web Framework gorilla mux and gohtml
* Build tool golang
* Testing tools testify
* Production FreshCloud on Google's Cloud Platform

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

Ensure postgres works locally.

```bash
"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
```

That's a wrap for now.
