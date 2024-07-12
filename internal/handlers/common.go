package handlers

import (
	"database/sql"
	"net/http"
	postgres "vortex/internal/db/postgre"
)

type ServiceHandler interface {
	AddClient(w http.ResponseWriter, r *http.Request)
	UpdateClient(w http.ResponseWriter, r *http.Request)
	DeleteClient(w http.ResponseWriter, r *http.Request)
	UpdateAlgorithmStatus(w http.ResponseWriter, r *http.Request)
}

type Service struct {
	DB *postgres.PostgresDriver
}

func NewService(pool *sql.DB) *Service {
	return &Service{
		DB: postgres.NewPostgresDriver(pool),
	}
}
