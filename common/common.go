package common

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/guregu/null.v4"
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

func ParseNullStringToTime(s null.String) (t null.Time) {
	if !s.Valid {
		return
	}

	ts, err := time.Parse(time.RFC3339, s.String)

	if err != nil {
		return
	}

	return null.TimeFrom(ts)
}
