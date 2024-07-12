package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vortex/internal/model"
)

// The AddClient method is a handler that reads the body of the POST request and
// passes its contents to methods which created new "client" and "algorithm" records.
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
