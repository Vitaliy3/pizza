package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type HttpError struct {
	Date  time.Time `json:"date"`
	Error string    `json:"message"`
}

func NewHttpError(w http.ResponseWriter, err error) []byte {
	w.WriteHeader(400)
	httpErr := HttpError{Error: err.Error()}
	data, _ := json.Marshal(httpErr)
	return data
}
