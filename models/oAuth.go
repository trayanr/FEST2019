package models

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetConfig() *oauth2.Config {
	clientId := "1060979839735-cajl2nl7sis3g2qsa1chba22liu8smod.apps.googleusercontent.com"
	clientSecret := "UDRVarko10cfVZYUy17Xs8Km"
	RedirectURL := "http://localhost:3000/api/oauth"

	conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/fitness.activity.write",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
