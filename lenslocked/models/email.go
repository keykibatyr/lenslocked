package models

import (
	"fmt"

	"github.com/go-mail/mail/v2"
)

const (
	DefaultSender = "keykibatyr@gmail.com"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type EmailService struct {
	DefualtSender string

	dialer *mail.Dialer
}

type Email struct {
	From string
	To string
	Subject string
	Plaintext string
	HTML string
}

func NewEmailService(config SMTPConfig) *EmailService {
	es := EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}

	return &es
}

func (es *EmailService) Send(email Email) error {
	msg := mail.NewMessage()
	msg.SetHeader("To", email.To)
	es.setForm(msg, email)
	msg.SetHeader("Subject", email.Subject)
	switch{
	case email.Plaintext != "" && email.HTML != "":
		msg.SetBody("text/plain", email.Plaintext)
		msg.AddAlternative("text/html", email.HTML)
	case email.Plaintext != "":
		msg.SetBody("text/plain", email.Plaintext)
	case email.HTML != "":
		msg.AddAlternative("text/HTML", email.HTML)
	}
	err := es.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("message could not send")
	}

	return nil
}

func (es *EmailService) setForm(msg *mail.Message, email Email) {
	var from string 
	switch {
	case email.From != "":
		from = email.From
	case es.DefualtSender != "":
		from = es.DefualtSender
	default:
		from = DefaultSender
	}
	msg.SetHeader("From", from)
}

func (es *EmailService) ForgotPassword(to, resetURL string) error {
	email := Email{
		To: to,
		Subject: "Reset your Password",
		Plaintext: "reset your password by visiting this url: " + resetURL,
		HTML: `<p>to reset your password please visit the following link: " <a href="` + resetURL + `">` + resetURL + `</a></p>`,
	}

	err := es.Send(email)
	if err != nil {
		return fmt.Errorf("could not send the passwoed reset link")
	}

	return nil
}

