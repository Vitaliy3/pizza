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
	mux.HandleFunc("/signin/", handlers.SignIn)
	mux.HandleFunc("/signup/", handlers.SignUp)
	mux.HandleFunc("/images/", handlers.Public)
	mux.HandleFunc("/items/", handlers.Items)
	mux.HandleFunc("/buy/", handlers.Buy)
	mux.HandleFunc("/category/", handlers.Category)
	mux.HandleFunc("/ban/", handlers.Ban)
	mux.HandleFunc("/phones/", handlers.Phones)
	mux.HandleFunc("/setrole/", handlers.SetRole)

	if err := http.ListenAndServe(":4500", mux); err != nil {
		log.Fatal(err)
	}
}
