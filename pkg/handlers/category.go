package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Category(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	category := models.Category{}
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("category err read body")
			return
		}
		err = json.Unmarshal(body, &category)
		if err != nil {
			fmt.Println("category err unmarshal")
			return
		}
		category.Save()
	}

	categories, _ := category.GetAll()
	b, _ := json.Marshal(categories)
	w.Write(b)
}
