package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"vortex/internal/model"
)

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	log.Print("UpdateClient:\t")
	var client model.Client
	var found = false
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println("err during encoding body: ", err)
	}
	for i := range Clients {
		if Clients[i].ID == client.ID {
			Clients[i].ClientName = client.ClientName
			Clients[i].Version = client.Version
			Clients[i].Image = client.Image
			Clients[i].CPU = client.CPU
			Clients[i].Memory = client.Memory
			Clients[i].Priority = client.Priority
			Clients[i].NeedRestart = client.NeedRestart
			Clients[i].UpdatedAt = time.Now()
			found = true
		}
	}
	if found == false {
		return
	}
	for i := range Clients {
		log.Println(Clients[i])
	}
}
