package common

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"gopkg.in/guregu/null.v4"
)

var (
	ErrNameMustAlphanumeric = errors.New("name must be alphanumeric")
	ErrNameTooShort         = errors.New("name must be above 3 character")
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

func ValidateName(name string) error {
	if len(name) < 3 {
		return ErrNameTooShort
	}
	if ok := strings.ContainsAny(name, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ .-_()[]{};:'\""); !ok {
		return ErrNameMustAlphanumeric
	}

	return nil
}
