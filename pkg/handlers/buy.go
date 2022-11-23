package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Buy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "POST" {
		urlToken := r.URL.Query().Get("token")
		session := models.Sessions[urlToken]

		userModel := models.User{Id: session.Id}
		user, _ := userModel.Select()
		if user.Blocked {
			w.Write([]byte("user blocked"))
			return
		}

		buy := models.Buy{}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("buy err readall:", err)
		}

		err = json.Unmarshal(data, &buy)
		if err != nil {
			fmt.Println("err unmarshal:", err)
		}

		err = buy.Buy(session.Id)
		if err != nil {
			fmt.Println("buy err handler:", err)
		}
		w.Write(nil)
	}
}
