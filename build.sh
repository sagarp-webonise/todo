#/bin/bash

goose up

xo pgsql://local:local@localhost/todo?sslmode=disable -o app/exp --suffix=.go --template-path templates/
