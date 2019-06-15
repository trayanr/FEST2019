package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/trayanr/FEST2019/drivers"
	"github.com/trayanr/FEST2019/models"
)

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

func Login(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	body := models.User{}
	decoder.Decode(&body)
	err := body.HashPassword()
	if err != nil {
		//
		log.Println("hash error")

	}
	user, err := drivers.GetUserByCredentials(body.Username, body.Password)
	if err != nil {
		//
		log.Println("driver err", err)
	}
	session, err := store.Get(r, "Auth")
	if err != nil {
		//
		log.Println("store err", err)
	}
	user, err = drivers.GetUserByCredentials(body.Username, body.Password)
	if err != nil {
		//
		log.Println("session err", err)
	}
	session.Values["auth"] = true
	session.Values["id"] = user.ID
	err = session.Save(r, w)
	if err != nil {
		log.Println("session save err", err)
	}

}
func Register(w http.ResponseWriter, r *http.Request) {

}

func RegisterAndLogin(w http.ResponseWriter, r *http.Request) {

}
