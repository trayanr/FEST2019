package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/trayanr/FEST2019/drivers"
	"github.com/trayanr/FEST2019/models"
	"golang.org/x/oauth2"
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

func oAuthCallback(w http.ResponseWriter, r *http.Request) {

	authCode := r.URL.Query().Get("code")
	log.Println(authCode)

	conf := models.GetConfig()

	tok, err := conf.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("token", tok)

	req, err := http.NewRequest("GET", "https://www.googleapis.com/fitness/v1/users/me/dataSources", nil)
	log.Println(err)
	req.Header.Set("Authorization", tok.AccessToken)
	client := conf.Client(oauth2.NoContext, tok)

	res, err := client.Do(req)
	log.Println(err)
	body, err := ioutil.ReadAll(res.Body)
	dataSources := map[string][]DataSource{
		"dataSource": []DataSource{},
	}
	err = json.Unmarshal(body, &dataSources)
	ds := dataSources["dataSource"]
	log.Println(err, len(ds))
	fmt.Println(ds)
	for _, d := range ds {
		GetForThisDay(d, tok)
	}
}

func GetConfigURL() ConfigURL {

	conf := models.GetConfig()
	url := conf.AuthCodeURL("state")

	res := ConfigURL{
		URL: url,
	}

	return res
}

type ConfigURL struct {
	URL string `json:"url"`
}
