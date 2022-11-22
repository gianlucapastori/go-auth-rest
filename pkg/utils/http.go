package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gianlucapastori/go-auth-jwt/pkg/validator"
)

func Respond(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func ReadRequest(req *http.Request, s interface{}) error {
	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&s); err != nil {
		return fmt.Errorf("could not unmarsh request body into struct: %v", err)
	}

	if err := validator.ValidateStruct(s); err != nil {
		return err
	}

	return nil
}
