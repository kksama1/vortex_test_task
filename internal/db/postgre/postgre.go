package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
	"time"
	"vortex/internal/model"
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
func createConnection() *sql.DB {

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

func SetUpDB() {
	db := createConnection()
	defer db.Close()
	sqlFile, err := os.Open("/usr/local/src/db/sql/client.sql")
	if err != nil {
		panic(err)
	}
	defer sqlFile.Close()

	sqlBytes, err := io.ReadAll(sqlFile)
	if err != nil {
		panic(err)
	}

	createTableQuery := string(sqlBytes)

	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	sqlFile, err = os.Open("/usr/local/src/db/sql/algorithm.sql")
	if err != nil {
		panic(err)
	}

	sqlBytes, err = io.ReadAll(sqlFile)
	if err != nil {
		panic(err)
	}

	createTableQuery = string(sqlBytes)

	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}
}

func GetTables() {
	db := createConnection()
	defer db.Close()
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	// Чтение результатов запроса
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			panic(err)
		}
		log.Println(tableName)
	}

}

func AddClient(client *model.Client) error {
	db := createConnection()
	defer db.Close()
	//query := `INSERT INTO clients(clientName, version, image, cpu, memory, priority, needRestart,
	//                spawnedAt, createdAt, updatedAt)
	//values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	query := `
	WITH inserted_client AS (
    INSERT INTO clients(clientName, version, image, cpu, memory, priority, needRestart)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id
	)
	INSERT INTO algorithm_status(clientID)
	SELECT id AS clientID
	FROM inserted_client;
`

	_, err := db.Query(query, client.ClientName, client.Version, client.Image, client.CPU, client.Memory,
		client.Priority, client.NeedRestart)
	if err != nil {
		return fmt.Errorf("error inserting client: %v", err)
	}

	return nil
}

func GetAllClients() ([]model.Client, error) {
	log.Println("postges.GetAllClients")
	db := createConnection()
	defer db.Close()
	var clients []model.Client
	rows, err := db.Query("SELECT * FROM clients")
	if err != nil {
		log.Println("SELECT ERR")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var client model.Client
		err = rows.Scan(&client.ID, &client.ClientName, &client.Version, &client.Image, &client.CPU, &client.Memory,
			&client.Priority, &client.NeedRestart, &client.SpawnedAt, &client.CreatedAt, &client.UpdatedAt)
		if err != nil {
			return nil, err
		}
		//log.Println("Client: ", client)
		clients = append(clients, client)
	}
	log.Println("Клиенты:", clients)
	return clients, nil
}

func GetAllAlgorithms() error {
	log.Println("GetAllAlgorithms")
	db := createConnection()
	defer db.Close()
	var algorithms []model.Algorithm

	rows, err := db.Query("SELECT * FROM algorithm_status")
	if err != nil {
		log.Println("SELECT ERR")
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var algorithm model.Algorithm
		err = rows.Scan(&algorithm.ID, &algorithm.ClientID, &algorithm.VWAP, &algorithm.TWAP, &algorithm.HFT)
		if err != nil {
			return err
		}
		log.Println("algorithm: ", algorithm)
		algorithms = append(algorithms, algorithm)
	}
	log.Println("Алгоритмы:", algorithms)
	return nil
}

func UpdateClient(client *model.Client) error {
	log.Println("UpdateClient")
	db := createConnection()
	defer db.Close()
	query := `UPDATE clients SET clientName=$1, version=$2, image=$3, cpu=$4,
                   memory=$5, priority=$6, needRestart=$7, updatedAt=$8 WHERE id=$9`
	log.Println("Updated")
	_, err := db.Exec(query,
		client.ClientName, client.Version, client.Image, client.CPU, client.Memory, client.Priority, client.NeedRestart,
		time.Now(), client.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteClient(client *model.Client) error {
	log.Println("DeleteClient")
	db := createConnection()
	defer db.Close()
	query := `
		DELETE FROM clients WHERE id = $1;
	`
	_, err := db.Exec(query, client.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DropAll() {
	db := createConnection()
	_, err := db.Exec("DROP TABLE algorithm_status")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = db.Exec("DROP TABLE clients")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Droped Tables")
}
