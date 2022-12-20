package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Buy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method == "POST" {
		buy := models.Buy{}
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Println("buy err readall:", err)
		}

		err = json.Unmarshal(data, &buy)
		if err != nil {
			fmt.Println("err unmarshal:", err)
		}

		session := models.User{}
		for _, k := range models.Sessions {
			if k.Telephone == buy.Phone {
				session = k
				return
			}
		}

		userModel := models.User{Id: session.Id}
		id, isBlocked, _ := userModel.CheckBan(buy.Phone)
		fmt.Println("isbl", isBlocked)
		fmt.Println("isbl", buy.Phone)
		if isBlocked {
			w.Write(NewHttpError(w, errors.New("Пользователь заблокирован")))
			return
		}
		fmt.Println("id:", id)
		err = buy.Buy(id)
		if err != nil {
			fmt.Println("buy err handler:", err)
		}
		w.Write(nil)
	}

	if r.Method == "PUT" {
		buy := models.Buy{}
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Println("buy err readall:", err)
		}

		err = json.Unmarshal(data, &buy)
		if err != nil {
			fmt.Println("err unmarshal:", err)
		}

		err = buy.StopTracking(buy)
		if err != nil {
			fmt.Println("=buy.StopTracking err:", err)
		}

		w.Write(nil)
		return
	}

	buy := models.Buy{}
	buyAll, _ := buy.BuyGetAll()
	bytes, _ := json.Marshal(buyAll)
	w.Write(bytes)
}
