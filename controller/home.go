package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gobuffalo/plush"
)

//GetLogin returns login page
func GetHome(w http.ResponseWriter, r *http.Request) {
	ctx := plush.NewContext()

	s, err := plush.Render(html("login.html"), ctx)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(s))
}

func html(file string) string {
	b, err := ioutil.ReadFile(fmt.Sprintf("./templates/%s", file))
	if err != nil {
		log.Fatal(err)
	}
	return string(b)

}
