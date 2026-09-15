package api

import (
	"encoding/json"
	"net/http"
)

// Server реализует ServerInterface - сгенерированный контракт из api/openapi.yaml.
type Server struct{}

func (s *Server) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Version{Version: "1.0.0"})
}

func (s *Server) CreateItem(w http.ResponseWriter, r *http.Request) {
	var req ItemCreate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Item{Id: 1, Title: req.Title})
}
