package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "postgres"
	port     = 5432
	database = "vortexTaskDB"
	username = "kksama"
	password = "kksama1"
)

type pgInfo struct {
	host     string
	port     int
	database string
	username string
	password string
}

// createConnection performs connection to db.
func CreateConnection() *sql.DB {

	ci := pgInfo{host: host, port: port, database: database, username: username, password: password}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		ci.host, ci.port, ci.username, ci.password, ci.database)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to postgres!")
	}
	return db
}
