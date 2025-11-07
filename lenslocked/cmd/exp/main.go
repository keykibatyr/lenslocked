package main

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "eaee543a90cd7e"
	password = "fdb7db38051211"
)

func main() {
	from := "keykibatyr@gmail.com"
	to := "alischreck678@gmail.com"
	subject := "This is my test email"
	plaintext := "This is the body of email"
	html := `<h1>HEllo there dude</h1><p>This is the eamil</p><p>Hope u enjoy it</p>`

	msg := mail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetHeader("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}

	fmt.Print("message sent")

}
