# todo-conversion-portal

# Setup
1. Setup postgresSQL on localhost.  
2. create a database named todo.The code assumed the user name and password to be `local` and the database name to be `todo`.
3. make a copy of `build.sh.example` as `build.sh`. Set the appropriate database params.  
4. Make a copy of `dbconf.yml.exampl` as `dbconf.yml` Set the appropriate database connection params.  



Migrations:
`goose up`

model generations:
`xo pgsql://<user>:<passowrd>@<host>/<db_name>?sslmode=disable -o internal/models --suffix=.go`


# To build:
`./build.sh`

# To run
`go run cmd/main.go`
