package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BanStruct struct {
	Phone string `json:"phone"`
}

func Ban(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		ban := BanStruct{}
		user := models.User{}
		data, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(data, &ban)
		fmt.Println("ban:", ban)

		err := user.Ban(ban.Phone)
		if err != nil {
			fmt.Println(err)
			w.Write(NewHttpError(w, err))
			return
		}
		w.Write(nil)
	}
}
