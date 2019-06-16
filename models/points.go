package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type Points struct {
	Value float64
	Level int
}

var types = map[int]float64{
	7:  1,
	8:  2.5,
	1:  1.8,
	82: 3,
}

var levels = map[int]float64{
	0: 0,
	1: 100,
	2: 300,
	3: 900,
	4: 2700,
	5: 8100,
}

func (p *Points) Calculate(activityType int, activityDuration uint) {
	var multiplier = types[activityType]
	act := activityDuration / 100000
	activityDuration = uint(act)
	var points float64
	if activityDuration != 0 {
		points = float64(activityDuration) * multiplier
	} else {
		points = multiplier
	}

	p.Value += points
	for i := 1; i <= len(levels); i++ {
		if p.Value < levels[i] && p.Value > levels[i-1] {
			p.Level = i
		}
	}
}

func GetLastSession(authCode string, lastTimeMS int64) ([]Session, error) {
	conf := GetConfig()

	token, err := conf.Exchange(oauth2.NoContext, "4/agF3GSPirhE6cqJWPJdFRBNuWRf0JljRzs-hTDmXKK5IfZAIVe4X1K9LP7gXMxwkiEIrorF6kFo4uR6fVeDKA08")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("token", token)

	now := time.Now()
	lastTime := time.Unix(0, lastTimeMS*int64(time.Millisecond))
	url := fmt.Sprint("https://www.googleapis.com/fitness/v1/users/me/sessions?startTime=", lastTime.Format("2006-01-02T15:04:05Z"), "&endTime=", now.Format("2006-01-02T15:04:05Z"))
	cl := conf.Client(oauth2.NoContext, token)

	req, err := http.NewRequest("GET", url, nil)
	log.Println(err)
	req.Header.Set("Authorization", token.AccessToken)
	res, err := cl.Do(req)
	log.Println(err)
	b, err := ioutil.ReadAll(res.Body)
	sess := map[string][]Session{}
	json.Unmarshal(b, &sess)
	fmt.Println(err)
	fmt.Println(string(b))
	return sess["session"], nil
}

//TODO::da pazim posledoto dobavqne na to4ki

func PointsTest() {
	sess, err := GetLastSession("4/agE8C2wUsc4ory3uHYdFyZC6JDc-owmwW9Ohjf8l65AaSKWBGLjFY2zaw4q0Y_Ural0kiD8fEkhdRV62CdcEoO8", 1560634000000)
	log.Println(sess, err)
}

type Session struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	StartTimeMillis    int    `json:"startTimeMillis"`
	EndTimeMillis      int    `json:"endTimeMillis"`
	ModifiedTimeMillis int    `json:"modifiedTimeMillis"`
	ActivityType       int    `json:"activtyType"`
}
