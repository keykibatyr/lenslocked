package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/keykibatyr/lenslocked/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	smtp_host := os.Getenv("SMTP_HOST")
	smtp_port := os.Getenv("SMTP_PORT")
	smtp_username := os.Getenv("SMTP_USERNAME")
	smtp_password := os.Getenv("SMTP_PASSWORD")

	port, err := strconv.Atoi(smtp_port)
	if err != nil {
		panic(err)
	}

	es := models.NewEmailService(models.SMTPConfig{
		Host:     smtp_host,
		Port:     port,
		Username: smtp_username,
		Password: smtp_password,
	})

	email := models.Email{
		From:      "keykibatyr@gmail.com",
		To:        "alischreck678@gmail.com",
		Subject:   "This is my test email",
		Plaintext: "This is the body of email",
		HTML:      `<h1>Hello there dude</h1><p>This is the eamil</p><p>Hope u enjoy it</p>`,
	}

	err = es.Send(email)
	if err != nil {
		panic(err)
	}

	fmt.Print("message sent")
}
