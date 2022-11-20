package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		models.SaveImage(w, r)

		return
	}

	p := models.Product{}
	data, err := p.GetAll()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	}
	b, _ := json.Marshal(data)
	w.Write(b)

}
