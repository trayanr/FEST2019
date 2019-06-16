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
	renderer.Render(w, map[string]interface{}{})

}

//GetAwards връща страницата със награди за даден user
func GetAwards(w http.ResponseWriter, r *http.Request) {
	renderer := R.HTML("awards.html")
	renderer.Render(w, map[string]interface{}{})

}

//GetContests returns страницата със всички състезания
func GetContests(w http.ResponseWriter, r *http.Request) {
	renderer := R.HTML("contests.html")
	renderer.Render(w, map[string]interface{}{})
}

func GetForThisDay(ds DataSource, token *oauth2.Token) {
	now := time.Now()
	yesterday := time.Now().Add(-24*time.Hour + time.Minute)
	// fmt.Printf("%+v\n", ds)
	url := fmt.Sprint("https://www.googleapis.com/fitness/v1/users/me/sessions?startTime=", yesterday.Format("2006-01-02T15:04:05Z"), "&endTime=", now.Format("2006-01-02T15:04:05Z"))
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
