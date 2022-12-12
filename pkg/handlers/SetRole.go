package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Role struct {
	Phone string `json:"phone"`
	Role  int    `json:"role"`
}

func SetRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		user := models.User{}
		role := Role{}
		data, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(data, &role)
		fmt.Println("role:", role)
		err := user.SetRole(role.Phone, role.Role)
		if err != nil {
			fmt.Println(err)
			w.Write(NewHttpError(w, err))
		}

		w.Write(nil)
	}
}
