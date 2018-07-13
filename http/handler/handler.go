package handler

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func writeError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&errorResponse{
		Error:  err.Error(),
		Status: status,
	})
}

func writeSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
