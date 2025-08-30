package api

import (
	"encoding/json"
	"net/http"
)

func WriteToJSON[T any](w http.ResponseWriter, payload T, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to encode response"}`))
	}
}
