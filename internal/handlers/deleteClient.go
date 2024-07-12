package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vortex/internal/model"
)

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
