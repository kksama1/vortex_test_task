package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"vortex/internal/model"
)

var Clients []model.Client

func AddClient(w http.ResponseWriter, r *http.Request) {
	log.Print("AddClient:\t")
	var client model.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println("err during encoding body: ", err)
	}
	client.CreatedAt = time.Now()
	client.SpawnedAt = time.Now()
	client.UpdatedAt = time.Now()
	Clients = append(Clients, client)

	for i := range Clients {
		log.Println(Clients[i])
	}

}
