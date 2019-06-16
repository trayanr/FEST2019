package models

import (
	"log"
	"testing"
)

func TestPoints(t *testing.T) {
	sess, err := GetLastSession("4/agFXxyLx1qvjdUsxAFBsEHQ9WBHiuC_1y4bp48atic2tLCEZhEjipEAgEVyOkRhD5EeL3IBZwmqsRZnQ_p5xXl8", 1560630000000)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(sess)
}
