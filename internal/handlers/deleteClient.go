package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	postgres "vortex/internal/db/postgre"
	"vortex/internal/model"
)

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteClient:\t")
	var client model.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println("err during encoding body: ", err)
	}
	err := postgres.DeleteClient(&client)
	if err != nil {
		log.Println(err)
	}
	_, err = postgres.GetAllClients()
	if err != nil {
		log.Println(err)
	}
}
