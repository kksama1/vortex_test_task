package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	postgres "vortex/internal/db/postgre"
	"vortex/internal/model"
)

func UpdateAlgorithmStatus(w http.ResponseWriter, r *http.Request) {
	log.Print("AddClient:\t")
	var Algorithm model.Algorithm
	if err := json.NewDecoder(r.Body).Decode(&Algorithm); err != nil {
		log.Println("err during encoding body: ", err)
	}
	if err := postgres.UpdateAlgorithmStatus(Algorithm); err != nil {
		log.Println(err)
	}
	if err := postgres.GetAllAlgorithms(); err != nil {
		log.Println(err)
	}
}
