package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/smtp"

	"SenderS/env"
	"github.com/streadway/amqp"
)

type Message struct {
	Id        int
	User      string
	Email     string
	Operation string
}

func ReaderSender(ch <-chan amqp.Delivery) {
	var dat Message
	var body bytes.Buffer

	for ms := range ch {
		err := json.Unmarshal(ms.Body, &dat)
		if err != nil {
			panic(err)
		}
		err = ms.Ack(true)
		if err != nil {
			panic(err)
		}
		tmpl, err := template.ParseFiles("template/reg.html")
		if err != nil {
			panic(err)
		}

		body.Write([]byte(fmt.Sprintf("Subject: This is text about %s  \n%s\n\n", dat.Operation, env.MimeHeaders)))
		err = tmpl.Execute(&body, dat)
		if err != nil {
			panic(err)
		}

		to := []string{dat.Email}
		auth := smtp.PlainAuth("", env.From, env.Pass, env.SmtpServ)
		err = smtp.SendMail(env.SmtpServ+":"+env.SmtpPort, auth, env.From, to, body.Bytes())
		if err != nil {
			panic(err)
		}
		fmt.Println("Message send!")

	}
}
