package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
