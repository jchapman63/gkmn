package server

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

func StoreGameState(game *Game) {
	connectionStr := "postgres://postgres:mysecretpassword@localhost:5433/postgres?sslmode=disable"
	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	// create game table
	_, err = conn.Query("CREATE TABLE IF NOT EXISTS game (gameResult JSONB);")
	if err != nil {
		panic(err)
	}

	// create json object and insert
	gameJSON, err := json.Marshal(game)
	if err != nil {
		panic(err)
	}
	queryString := fmt.Sprintf("INSERT INTO game (gameResult)\nVALUES ('%s')", gameJSON)
	_, err = conn.Query(queryString)
	if err != nil {
		panic(err)
	}
}

func TestConnection() []string {
	// connect to pre created user and database (created to test out pgsql in a docker container)
	connectionStr := "postgres://postgres:mysecretpassword@localhost:5433/postgres?sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	// all results of all queries
	var results []string
	rows, err := conn.Query("SELECT version();")
	if err != nil {
		panic(err)
	}
	// iterates until there is not another row in the resulting table from the query
	for rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Println(version)
	}

	_, err = conn.Query("CREATE TABLE IF NOT EXISTS gkmn (firstData VARCHAR(255));")
	if err != nil {
		panic(err)
	}

	queryString := fmt.Sprintf("INSERT INTO gkmn (firstData)\nVALUES ('%s');", "NewHello")
	_, err = conn.Query(queryString)
	if err != nil {
		panic(err)
	}
	rows, err = conn.Query("SELECT * FROM gkmn;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var result string
		rows.Scan(&result)
		results = append(results, result)
	}

	conn.Close()

	return results
}

// connection string format
// postgres://user:password@host/database
