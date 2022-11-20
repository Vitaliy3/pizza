package main

import (
	"agile/pkg/dbManager"
	"agile/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	if err := dbManager.Init(); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/signIn/", handlers.SignIn)
	mux.HandleFunc("/signUp/", handlers.SignUp)
	mux.HandleFunc("/images/", handlers.Public)
	mux.HandleFunc("/items/", handlers.GetAll)

	if err := http.ListenAndServe(":4500", mux); err != nil {
		log.Fatal(err)
	}
}
