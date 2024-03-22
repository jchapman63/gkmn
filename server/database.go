package server

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func TestConnection() {
	// connect to pre created user and database (created to test out pgsql in a docker container)
	connectionStr := "postgres://docker_user:docker_user@localhost:5433/docker_user?sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	rows, err := conn.Query("SELECT version();")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Println(version)
	}

	conn.Close()
}

// connection string format
// postgres://user:password@host/database
