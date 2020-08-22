package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

const (
	// MIME Header
	MIME string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// MailSubject is the subject of the Mail
	MailSubject string = "🍃 Oops! Someone Stick Bugged you! 😭"
)

// Mail is the Mail to send
type Mail struct {
	From   string  // Email of the sender
	To     string  // Email of the Stick Bugged victim
	Body   string  // Body of the Mail
	Config *Config // Configuration required to authenticate and send the Mail
}

func (m *Mail) parseTemplate() error {
	tpl, err := template.ParseFiles("./email.html")

	if err != nil {
		return err
	}

	bfr := new(bytes.Buffer)

	if err = tpl.Execute(bfr, nil); err != nil {
		return err
	}

	m.Body = bfr.String()
	return nil
}

// Send the mail
func (m *Mail) Send() {
	// Create content (which includes some headers)
	if err := m.parseTemplate(); err != nil {
		log.Fatalln("Error while parsing content!\n" + err.Error())
	}
	ct := "To: " + m.To + "\r\nSubject: " + MailSubject + "\r\n" + MIME + "\r\n" + m.Body

	err := smtp.SendMail(fmt.Sprintf("%s:%d", m.Config.Host, m.Config.Port), smtp.PlainAuth("", m.Config.Email, m.Config.Password, m.Config.Host), m.Config.Email, []string{m.To}, []byte(ct))

	if err != nil {
		log.Fatalln("FAIL! Mail " + fmt.Sprintf("%s (%s) -> %s. Host: %s:%d!\n", m.From, m.Config.Email, m.To, m.Config.Host, m.Config.Port) + err.Error())
	} else {
		log.Fatalln("SUCCEED! Mail " + fmt.Sprintf("%s (%s) -> %s. Host: %s:%d!\n", m.From, m.Config.Email, m.To, m.Config.Host, m.Config.Port))
	}
}

// NewMail creates a new mail request
func NewMail(to string) *Mail {
	cf, err := GetConfig()

	if err != nil {
		log.Fatal("Error while getting the configuration file!\n" + err.Error())
	}

	return &Mail{
		From:   cf.Email,
		To:     to,
		Config: cf,
	}
}
