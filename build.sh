#/bin/bash

goose up

xo pgsql://local:local@localhost/todo?sslmode=disable -o app/domain --suffix=.go --template-path templates/
