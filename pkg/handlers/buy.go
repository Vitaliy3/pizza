package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Buy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method == "POST" {
		buy := models.Buy{}
		data, err := ioutil.ReadAll(r.Body)
		fmt.Println("data:", data)

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
		url, err := url.Parse(r.URL.String())
		if err != nil {
			fmt.Println("url.parse err:", err)
			return
		}

		params := strings.Split(url.String(), "/")
		buy := models.Buy{}
		fmt.Println(params[2])
		itemIdStr := strings.TrimSpace(params[2])
		itemId, _ := strconv.Atoi(itemIdStr)
		fmt.Println("STOP ", itemId)
		err = buy.StopTracking(itemId)
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
