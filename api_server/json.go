package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	type invalid struct {
		Error string `json:"error"`
	}

	invalidData := invalid{
		Error: msg,
	}

	respondWithJson(w, statusCode, invalidData)
}

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
