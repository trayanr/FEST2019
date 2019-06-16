package models

import (
	"log"
	"testing"
)

func TestPoints(t *testing.T) {
	sess, err := GetLastSession("4/agERWOUCHELUNVyJHpDeEUJnrLjsNwXvNmR0nCpz0cu0eebSxgpkKl2tOOxMnb6-qR-aP8C0HJZ4p3E0_wnEABo", 1560634000000)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(sess)
}
