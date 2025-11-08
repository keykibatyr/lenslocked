package main

import (
	"fmt"

	"github.com/keykibatyr/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "eaee543a90cd7e"
	password = "fdb7db38051211"
)

func main() {
	es := models.NewEmailService(models.SMTPConfig{
		Host: host,
		Port: port,
		Username: username,
		Password: password,
	})

	email := models.Email{
		From: "keykibatyr@gmail.com",
		To: "alischreck678@gmail.com",
		Subject: "This is my test email",
		Plaintext: "This is the body of email",
		HTML: `<h1>Hello there dude</h1><p>This is the eamil</p><p>Hope u enjoy it</p>`,
	}

	err := es.Send(email) 
	if err != nil {
		panic(err)
	}

	fmt.Print("message sent")
}
