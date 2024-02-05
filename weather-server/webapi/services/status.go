package services

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func StatusService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := StatusResponse{
			Status: "OK",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Println("JSON encode error in status service: ", err)
			return
		}
	}
}
