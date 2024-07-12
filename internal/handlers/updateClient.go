package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vortex/internal/model"
)

// UpdateClient method is a handler that reads the body of the POST request
// and passes its contents to method which modifies specified Client record.
func (s *Service) UpdateClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println("err during encoding body: ", err)
	}

	if err := s.DB.UpdateClient(&client); err != nil {
		log.Println(err)
		return
	}

	log.Println("Client updated")
}
