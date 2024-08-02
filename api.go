package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	service Service
}

func NewApiServer(s Service) *ApiServer {
	return &ApiServer{service: s}
}

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.service.GetCatFact(context.Background())
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, fact)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
