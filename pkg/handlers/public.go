package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func Public(w http.ResponseWriter, r *http.Request) {
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
		file, err := os.ReadFile("uploads/" + imageName)
		if err != nil {
			log.Println(err)
			w.Write([]byte("File not found"))
		}

		w.Write(file)
	}
}
