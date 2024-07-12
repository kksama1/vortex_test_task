package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vortex/internal/model"
)

func (s *Service) AddClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println("err during encoding body: ", err)
		return
	}

	if err := s.DB.AddClient(&client); err != nil {
		log.Println(err)
		return
	}

	log.Println("Client added")

}
