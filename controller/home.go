package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/trayanr/FEST2019/models"
	"golang.org/x/oauth2"
)

//GetLogin returns welcome page
func GetWelcome(w http.ResponseWriter, r *http.Request) {
	renderer := R.HTML("welcome.html", "layouts/main.html")
	renderer.Render(w, map[string]interface{}{})
}

//GetLogin returns home page (на логнат потребител)
func GetHome(w http.ResponseWriter, r *http.Request) {
	renderer := R.HTML("home.html")
	renderer.Render(w, map[string]interface{}{})
}

//GetLogin returns home page (на логнат потребител)
func GetProfile(w http.ResponseWriter, r *http.Request) {
	renderer := R.HTML("profile.html")
	yo := renderer.Render(w, map[string]interface{}{})
	fmt.Println(yo.Error())
}

func GetForThisDay(ds DataSource, token *oauth2.Token) {
	now := time.Now()
	yesterday := time.Now().Add(-24*time.Hour + time.Minute)
	// fmt.Printf("%+v\n", ds)
	url := fmt.Sprint("https://www.googleapis.com/fitness/v1/users/me/dataSources/", ds.DataStreamName, "/datasets/", yesterday.UnixNano(), "-", now.UnixNano())
	// fmt.Println(url)
	conf := models.GetConfig()
	cl := conf.Client(oauth2.NoContext, token)

	req, err := http.NewRequest("GET", url, nil)
	log.Println(err)
	req.Header.Set("Authorization", token.AccessToken)
	res, err := cl.Do(req)
	log.Println(err)
	b, err := ioutil.ReadAll(res.Body)
	fmt.Println(err)
	fmt.Println(string(b))
}

type DataSource struct {
	DataStreamID        string        `json:"dataStreamId"`
	DataStreamName      string        `json:"dataStreamName"`
	Type                string        `json:"type"`
	DataType            DataType      `json:"dataType"`
	Application         Application   `json:"application"`
	DataQualityStandard []interface{} `json:"dataQualityStandard"`
	Device              Device        `json:"device"`
}

type DataType struct {
	Name   string      `json:"name"`
	Fields []DataField `json:"field"`
}

type Application struct {
	PackageName string `json:"packageName"`
	Version     string `json:"version"`
}

type DataField struct {
	Name   string `json:"duration"`
	Format string `json:"format"`
}

type Device struct {
	UID          string `json:"uid"`
	Type         string `json:"phone"`
	Version      string `json:"verison"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
}

func init() {
	conf := models.GetConfig()
	url := conf.AuthCodeURL("state")
	log.Println(url)
}
