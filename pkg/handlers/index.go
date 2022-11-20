package handlers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}
