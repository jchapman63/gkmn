package server

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func TestConnection() {
	// connect to pre created user and database (created to test out pgsql in a docker container)
	connectionStr := "postgres://docker_user/docker_user@localhost:5433/docker_user"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	conn.Close()
}

// connection string format
// postgres://user:password@host/database
