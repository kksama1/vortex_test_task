package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vortex/internal/model"
)

func (s *Service) UpdateAlgorithmStatus(w http.ResponseWriter, r *http.Request) {
	var Algorithm model.Algorithm

	if err := json.NewDecoder(r.Body).Decode(&Algorithm); err != nil {
		log.Println("err during encoding body: ", err)
		return
	}

	if err := s.DB.UpdateAlgorithmStatus(Algorithm); err != nil {
		log.Println(err)
		return
	}

	log.Println("Algorithm status updated")
}
