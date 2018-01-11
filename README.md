# todo-conversion-portal

Migrations:
`goose up`

model generations:
`xo pgsql://local:local@localhost/todo?sslmode=disable -o internal/models --suffix=.go`


#To build:
`./build.sh`

#To run
`go run cmd/main.go`
