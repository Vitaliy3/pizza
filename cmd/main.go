package main

import (
	"agile/pkg/dbManager"
	"agile/pkg/models"
	"agile/pkg/models/auth"
	"agile/pkg/models/session"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
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

func publicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/public")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmp := strings.Split(r.URL.String(), "/")
	fmt.Println(r.URL.String())
	imageName := tmp[len(tmp)-1]
	fmt.Println("imageName:", imageName)
	if imageName != "" {
		file, err := os.ReadFile("../uploads/" + imageName)
		if err != nil {
			log.Println(err)
			w.Write([]byte("File not found"))
		}

		w.Write(file)
	}
}

func signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		user := auth.User{}
		data, err := ioutil.ReadAll(r.Body)
		fmt.Println("data", string(data))
		if err != nil {
			w.Write([]byte(err.Error()))
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
			w.Write([]byte(err.Error()))
			return
		}

		token := session.Add(user)
		data, err = json.Marshal(token)
		if err != nil {
			fmt.Println("err marshal token:", err)
		}

		w.Write(data)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		user := auth.User{}
		data, err := ioutil.ReadAll(r.Body)
		fmt.Println("data", string(data))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		err = json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println("signin: err unmarshall user ", err)
			return
		}

		err = user.SignUp()
		if err != nil {
			fmt.Println("SignUp err: ", err)
			return
		}

		err = user.SignIn()
		if err != nil {
			fmt.Println("user.SignIn err:", err)
			w.Write([]byte(err.Error()))
			return
		}

		token := session.Add(user)
		data, err = json.Marshal(token)
		if err != nil {
			fmt.Println("err marshal token:", err)
		}

		w.Write(data)

	}
}

func main() {
	dbManager.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/images/", publicHandler)
	mux.HandleFunc("/items/", getAllHandler)
	mux.HandleFunc("/signin/", signin)
	mux.HandleFunc("/signup/", signup)

	if err := http.ListenAndServe(":4500", mux); err != nil {
		log.Fatal(err)
	}
}
