package gomail

import (
	"log"
	"sendMail/config/env"
	"testing"
	"time"
)

func TestSendMail(t *testing.T) {
	env.Init("./config.json")
	c := make(chan error)
	body := `
	Hello there,
	This is a test message from Ghoghnos travel agency,
	Please Dont reply ;)
	
	your faithfully, ghoghnos
`
	go SendMail("ghoroubi85@gmail.com", "test", body, c)
	time.Sleep(time.Second * 2)
	if err := <-c; err != nil {
		log.Println(err)
		t.FailNow()
	}
}
