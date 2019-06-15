package main

import (
	"net/http"

	"github.com/trayanr/FEST2019/controller"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", controller.Login)
	mux.Handle("/static", http.FileServer(http.Dir("static")))

	// mux.HandleFunc("/", controller.GetHome) //home page
	http.ListenAndServe(":3000", mux)
}
