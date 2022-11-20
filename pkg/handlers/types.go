package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type HttpError struct {
	Date  time.Time `json:"date"`
	Error string    `json:"error"`
}

func NewHttpError(w http.ResponseWriter, err error) []byte {
	w.WriteHeader(400)
	data, _ := json.Marshal(HttpError{Date: time.Now(), Error: err.Error()})
	return data
}
