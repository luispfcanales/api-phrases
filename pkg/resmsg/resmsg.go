package resmsg

import (
	"encoding/json"
	"net/http"
)

// ResponseMSG is a Template to response body
type ResponseMSG struct {
	Status  int
	Message string
	Data    interface{}
}

// New Create and return a message to response body
func New() *ResponseMSG {
	return &ResponseMSG{
		Status:  http.StatusOK,
		Message: "success",
	}
}

// SendReponse send message in format JSON
func (r *ResponseMSG) SendReponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(&r)
}
