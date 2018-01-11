#/bin/bash

goose up

xo pgsql://local:local@localhost/todo?sslmode=disable -o internal/models --suffix=.go --template-path templates/
