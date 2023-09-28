#!/bin/bash -eu

docker run --rm -p 5432:5432 --name postgres-smak -e POSTGRES_PASSWORD=password  postgres
