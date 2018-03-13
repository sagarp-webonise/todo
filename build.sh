#/bin/bash

goose -dir="./db/migrations" postgres "user=local password=local dbname=local sslmode=disable" up

xo pgsql://local:local@localhost/local?sslmode=disable -o app/models --suffix=.go --template-path templates/
