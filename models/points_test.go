package models

import (
	"log"
	"testing"
)

func TestPoints(t *testing.T) {
	sess, err := GetLastSession("4/agE8C2wUsc4ory3uHYdFyZC6JDc-owmwW9Ohjf8l65AaSKWBGLjFY2zaw4q0Y_Ural0kiD8fEkhdRV62CdcEoO8", 1560634000000)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(sess)
}
