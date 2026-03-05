package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrMissingBody = errors.New("missing request body")

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return ErrMissingBody
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(v)
}
