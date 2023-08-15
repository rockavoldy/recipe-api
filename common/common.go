package common

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func WriteResponse(w http.ResponseWriter, status int, resp Response) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(resp)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	resp := Response{
		Message: err.Error(),
		Data:    nil,
	}
	WriteResponse(w, status, resp)
}
