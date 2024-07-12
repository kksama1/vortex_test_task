package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vortex/internal/model"
)

// The DeleteClient method is a handler that reads the body of the POST request
// and passes its contents to method which deletes specified "client" and
// "algorithm" records.
func (s *Service) DeleteClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println("err during encoding body: ", err)
		return
	}

	err := s.DB.DeleteClient(&client)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Сlient deleted")
}
