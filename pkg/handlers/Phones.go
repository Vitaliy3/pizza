package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Phones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "GET" {
		user := models.User{}
		users, err := user.GetAll()
		if err != nil {
			fmt.Println(err)
			return
		}

		phones := make([]string, 0)
		for _, v := range users {
			phones = append(phones, v.Telephone)
		}

		b, _ := json.Marshal(phones)
		w.Write(b)
	}
}
