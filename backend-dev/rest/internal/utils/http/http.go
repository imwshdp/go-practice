package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	_ = WriteJSON(w, status, map[string]string{
		"error": err.Error(),
	})
}
