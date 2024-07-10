package postgre

import (
	"database/sql"
	"fmt"
)

type pgInfo struct {
	host     string
	port     int
	database string
	username string
	password string
}

// createConnection performs connection to db.
func createConnection() *sql.DB {

	ci := pgInfo{host: "localhost", port: 5432, database: "testing", username: "admin_pg", password: "root"}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		ci.host, ci.port, ci.username, ci.password, ci.database)
	db, err := sql.Open("postgre", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
