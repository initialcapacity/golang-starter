#!/bin/bash
set -e
apk add postgresql-client postgresql-libs
while ! pg_isready --user starter --host postgresql &> /dev/null; do
  sleep 2
  echo "Waiting for the starter database to become active."
done
migrate -verbose -path '/home/databases/starter' -database 'postgres://starter:starter@postgresql:5432/starter_development?sslmode=disable' up
