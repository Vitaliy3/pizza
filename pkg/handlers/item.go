package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Items(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "POST" {
		models.SaveImage(w, r)
		return
	}

	product := models.Product{}
	data, err := product.GetAll()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	}

	b, _ := json.Marshal(data)
	w.Write(b)

}
