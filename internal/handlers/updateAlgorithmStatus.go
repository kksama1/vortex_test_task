package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vortex/internal/model"
)

// UpdateAlgorithmStatus method is a handler that reads the body of the POST request
// and passes its contents to method which modifies specified Algorithm.
func (s *Service) UpdateAlgorithmStatus(w http.ResponseWriter, r *http.Request) {
	var Algorithm model.Algorithm

	if err := json.NewDecoder(r.Body).Decode(&Algorithm); err != nil {
		log.Println("err during encoding body: ", err)
		return
	}

	if err := s.DB.UpdateAlgorithmStatus(&Algorithm); err != nil {
		log.Println(err)
		return
	}

	log.Println("Algorithm status updated")
}
