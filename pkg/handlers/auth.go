package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		user := models.User{}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write(NewHttpError(w, err))
			return
		}

		err = json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println("signin: err unmarshall user ", err)
			return
		}

		err = user.SignIn()
		if err != nil {
			fmt.Println("user.SignIn err:", err)
			w.Write(NewHttpError(w, errors.New("Неверный логин или пароль!")))
			return
		}

		token := models.Add(user)
		data, err = json.Marshal(token)
		if err != nil {
			w.Write(NewHttpError(w, err))
			fmt.Println("err marshal token:", err)
		}

		w.Write(data)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		user := models.User{}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write(NewHttpError(w, err))
			return
		}

		err = json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println("signin: err unmarshall user ", err)
			w.Write(NewHttpError(w, err))
			return
		}

		err = user.SignUp()
		if err != nil {
			w.Write(NewHttpError(w, err))
			return
		}

		err = user.SignIn()
		if err != nil {
			fmt.Println("user.SignIn err:", err)
			w.Write(NewHttpError(w, err))
			return
		}

		token := models.Add(user)
		data, err = json.Marshal(token)
		if err != nil {
			fmt.Println("err marshal token:", err)
			w.Write(NewHttpError(w, err))
		}

		w.Write(data)

	}
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		user := models.User{}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write(NewHttpError(w, err))
			return
		}

		err = json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println("signin: err unmarshall user ", err)
			w.Write(NewHttpError(w, err))
			return
		}
		fmt.Println("user pidor", user)
		err = user.Update()
		if err != nil {
			w.Write(NewHttpError(w, err))
			return
		}
	}
}
