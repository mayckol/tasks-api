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

// JSONSingleResPresenter is a dynamic presenter for returning JSON single response.
func JSONSingleResPresenter(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if encodeErr := json.NewEncoder(w).Encode(data); encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

type JSONError400Response struct {
	Data  []validation.InvalidFields `json:"data"`
	Error string                     `json:"error"`
}

type ListResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Page  int         `json:"page,omitempty"`
	Total int         `json:"total,omitempty"`
}

// JSONListPresenter is a dynamic presenter for returning list JSON response.
func JSONListPresenter(w http.ResponseWriter, res ListResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if encodeErr := json.NewEncoder(w).Encode(res); encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
