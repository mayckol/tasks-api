package presenter

import (
	"encoding/json"
	"net/http"
	"tasks-api/internal/validation"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

// JSONPresenter is a dynamic presenter for returning JSON responses.
func JSONPresenter(w http.ResponseWriter, statusCode int, data interface{}, err ...error) {
	var vErr error
	if len(err) > 0 {
		vErr = err[0]
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{}
	if vErr != nil {
		response.Error = vErr.Error()
	}

	if data != nil {
		response.Data = data
	}

	if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

type JSONError400Response struct {
	Data  []validation.InvalidFields `json:"data"`
	Error string                     `json:"error"`
}
