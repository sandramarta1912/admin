#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o admin .

MYSQL_HOST=localhost MYSQL_PORT=3306 MYSQL_ROOT=root MYSQL_ROOT_PASSWORD=cms MYSQL_DATABASE=adserver ./admin