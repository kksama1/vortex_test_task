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

type DatabaseDriver interface {
	//CreateConnection(host string, port int, database, username, password string)
	SetUpDB()
	//GetTables()
	AddClient(client *model.Client) error
	GetAllClients() ([]model.Client, error)
	GetAllAlgorithms() error
	UpdateClient(client *model.Client) error
	UpdateAlgorithmStatus(algorithm model.Algorithm) error
	DeleteClient(client *model.Client) error
	GetActiveAlgorithms() ([]model.Algorithm, error)
	GetInActiveAlgorithms() ([]model.Algorithm, error)
	//DropAll()
	CloseConnection() error
}

type PostgresDriver struct {
	Pool *sql.DB
}

func NewPostgresDriver(pool *sql.DB) *PostgresDriver {
	return &PostgresDriver{Pool: pool}
}

// createConnection performs connection to Db.

func CreateConnection(
	host string,
	port int,
	database string,
	username string,
	password string,
	sslmode string,
) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, username, password, database, sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// Установка максимального количества открытых соединений
	db.SetMaxOpenConns(15)
	// Установка максимального количества соединений в пуле
	db.SetMaxIdleConns(5)
	// Установка максимального времени жизни соединения в пуле
	db.SetConnMaxLifetime(time.Minute * 5)

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to postgres!")
	}
	return db
}

func (p *PostgresDriver) SetUpDB() {
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

	_, err = p.Pool.Exec(createTableQuery)
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

	_, err = p.Pool.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}
}

func (p *PostgresDriver) GetTables() {
	rows, err := p.Pool.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	// Чтение результатов запроса
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (p *PostgresDriver) AddClient(client *model.Client) error {
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
	_, err := p.Pool.Exec(query, client.ClientName, client.Version, client.Image, client.CPU, client.Memory,
		client.Priority, client.NeedRestart)
	if err != nil {
		return fmt.Errorf("error inserting client: %v", err)
	}
	return nil
}

func (p *PostgresDriver) GetAllClients() ([]model.Client, error) {
	var clients []model.Client
	rows, err := p.Pool.Query("SELECT * FROM clients")
	if err != nil {

		return nil, fmt.Errorf("error while selecting all clients: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var client model.Client
		err = rows.Scan(&client.ID, &client.ClientName, &client.Version, &client.Image, &client.CPU, &client.Memory,
			&client.Priority, &client.NeedRestart, &client.SpawnedAt, &client.CreatedAt, &client.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error while scanning rows: %v", err)
		}
		clients = append(clients, client)
	}
	log.Println("Клиенты:", clients)
	return clients, nil
}

func (p *PostgresDriver) GetAllAlgorithms() error {
	var algorithms []model.Algorithm

	rows, err := p.Pool.Query("SELECT * FROM algorithm_status")
	if err != nil {
		return fmt.Errorf("error while selecting all algorithms: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var algorithm model.Algorithm
		err = rows.Scan(&algorithm.ID, &algorithm.ClientID, &algorithm.VWAP, &algorithm.TWAP, &algorithm.HFT)
		if err != nil {
			return fmt.Errorf("error while scanning rows: %v", err)
		}
		algorithms = append(algorithms, algorithm)
	}
	return nil
}

func (p *PostgresDriver) UpdateClient(client *model.Client) error {

	query := `UPDATE clients SET clientName=$1, version=$2, image=$3, cpu=$4,
                   memory=$5, priority=$6, needRestart=$7, updatedAt=$8 WHERE id=$9`

	_, err := p.Pool.Exec(query,
		client.ClientName, client.Version, client.Image, client.CPU, client.Memory, client.Priority, client.NeedRestart,
		time.Now(), client.ID)
	if err != nil {
		return fmt.Errorf("error while scanning rows: %v", err)
	}
	return nil
}

func (p *PostgresDriver) UpdateAlgorithmStatus(algorithm model.Algorithm) error {
	query := `
	UPDATE algorithm_status SET vwap=$1, twap=$2, hft=$3 WHERE clientID=$4
	`
	_, err := p.Pool.Exec(query,
		algorithm.VWAP, algorithm.TWAP, algorithm.HFT, algorithm.ClientID)
	if err != nil {
		return fmt.Errorf("error while updating algorithm status: %v", err)

	}
	return nil
}

func (p *PostgresDriver) DeleteClient(client *model.Client) error {
	query := `
		DELETE FROM clients WHERE id = $1;
	`
	_, err := p.Pool.Exec(query, client.ID)
	if err != nil {
		return fmt.Errorf("error while deleting client: %v", err)
	}
	return nil
}

func (p *PostgresDriver) GetActiveAlgorithms() ([]model.Algorithm, error) {
	query := `
		SELECT * FROM algorithm_status WHERE VWAP = TRUE OR TWAP = TRUE OR HFT = TRUE;
	`
	var algorithms []model.Algorithm
	rows, err := p.Pool.Query(query)
	if err != nil {
		log.Println("SELECT ERR")
		return nil, fmt.Errorf("error while selecting active algorithms: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var algorithm model.Algorithm
		err = rows.Scan(&algorithm.ID, &algorithm.ClientID, &algorithm.VWAP, &algorithm.TWAP, &algorithm.HFT)
		if err != nil {
			return nil, fmt.Errorf("error while scanning rows: %v", err)
		}
		log.Println("active algorithm: ", algorithm)
		algorithms = append(algorithms, algorithm)
	}
	return algorithms, nil
}

func (p *PostgresDriver) GetInActiveAlgorithms() ([]model.Algorithm, error) {
	log.Println("GetInActiveAlgorithms")
	query := `
		SELECT * FROM algorithm_status WHERE VWAP = FALSE AND TWAP = FALSE AND HFT = FALSE;
	`
	var algorithms []model.Algorithm
	rows, err := p.Pool.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error while selecting inactive algorithms: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var algorithm model.Algorithm
		err = rows.Scan(&algorithm.ID, &algorithm.ClientID, &algorithm.VWAP, &algorithm.TWAP, &algorithm.HFT)
		if err != nil {
			return nil, fmt.Errorf("error while scanning rows: %v", err)
		}
		log.Println("inactive algorithm: ", algorithm)
		algorithms = append(algorithms, algorithm)
	}
	return algorithms, nil
}

//func (p *PostgresDriver) DropAll() {
//	_, err := p.db.Exec("DROP TABLE algorithm_status")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	_, err = p.db.Exec("DROP TABLE clients")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//}

func (p *PostgresDriver) CloseConnection() error {
	err := p.Pool.Close()
	if err != nil {
		return fmt.Errorf("error while closing conection: %v", err)
	}
	log.Println("connection closed")
	return nil
}
