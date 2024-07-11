package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vortex/internal/model"
)

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteClient:\t")
	var NewClients []model.Client
	var client model.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println("err during encoding body: ", err)
	}
	for i := range Clients {
		if Clients[i].ID != client.ID {
			NewClients = append(NewClients, Clients[i])
		}
	}
	Clients = NewClients
	for i := range Clients {
		log.Println(Clients[i])
	}
}
