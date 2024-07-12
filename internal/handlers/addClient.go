package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	postgres "vortex/internal/db/postgre"
	"vortex/internal/model"
)

var Clients []model.Client

func AddClient(w http.ResponseWriter, r *http.Request) {
	log.Print("AddClient:\t")
	var clientToCreate model.Client
	if err := json.NewDecoder(r.Body).Decode(&clientToCreate); err != nil {
		log.Println("err during encoding body: ", err)
	}

	client := model.NewClient(clientToCreate)
	if err := postgres.AddClient(client); err != nil {
		log.Println(err)
	}
	log.Println("client added ro db")
	_, err := postgres.GetAllClients()
	if err != nil {
		log.Println(err)
	}
	//clients, _ := postgres.GetAllClients()
	log.Println(client)
	response, err := json.Marshal(client)
	if err != nil {
		log.Println("Marshal Err")
	}
	err = postgres.GetAllAlgorithms()
	if err != nil {
		log.Println("GetAllAlgorithms Err")
	}
	w.Write(response)

}
