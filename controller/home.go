package controllers

import (
	"net/http"
)

//GetLogin returns login page
func GetHome(w http.ResponseWriter, r *http.Request) {

	renderer := R.HTML("login.html")
	renderer.Render(w, map[string]interface{}{})

}
