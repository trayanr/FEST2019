package main

import (
	"fmt"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/trayanr/FEST2019/controller"
)

var port = os.Getenv("PORT")

func routes(r *mux.Router) {

	//ROUTE-ОВЕТЕ СЕ АДДВАТ ТУК :)

	addHandler(r, "/", controllers.GetWelcome).Methods("GET")
	addHandler(r, "/home", controllers.GetHome).Methods("GET")
	addHandler(r, "/profile", controllers.GetProfile).Methods("GET")
	addHandler(r, "/awards", controllers.GetAwards).Methods("GET")
	addHandler(r, "/contests", controllers.GetContests).Methods("GET")
	addHandler(r, "/api/login", controllers.Login).Methods("GET", "POST")
	addHandler(r, "/api/oauthPost", controllers.OAuthPost).Methods("POST", "GET")
	addHandler(r, "/api/oauth", controllers.OAuthCallback).Methods("GET", "POST")

}

func main() {

	r := mux.NewRouter()

	routes(r)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)
	fmt.Printf("listening on localhost:%s/ \n", port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
	if err != nil {
		log.Print(err.Error())
	}

}

func addHandler(r *mux.Router, path string, handler http.HandlerFunc) *mux.Route {
	return r.HandleFunc(path, handler) //.Host(conf.Subdomain + "." + conf.Host)
}

func startTimer() {
	// timer1 := time.

}
