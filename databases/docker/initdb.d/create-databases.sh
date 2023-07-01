#!/bin/bash
set -e
psql -v on_error_stop=1 --username postgresql <<-EOSQL
  create database starter_development;
  create user starter with password 'starter';
  grant all privileges on database starter_development to starter;
EOSQL
