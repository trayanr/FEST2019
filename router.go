package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/trayanr/FEST2019/controller"
	"net/http"
)

const port = 3000

func routes(r *mux.Router) {

	//ROUTE-ОВЕТЕ СЕ АДДВАТ ТУК :)

	addHandler(r, "/", controllers.GetHome).Methods("GET")
	addHandler(r, "/api/login", controllers.Login).Methods("GET", "POST")

}

func main() {

	r := mux.NewRouter()

	routes(r)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)
	fmt.Printf("listening on localhost:%d/ \n", port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	if err != nil {
		log.Print(err.Error())
	}

}

func addHandler(r *mux.Router, path string, handler http.HandlerFunc) *mux.Route {
	return r.HandleFunc(path, handler) //.Host(conf.Subdomain + "." + conf.Host)
}
