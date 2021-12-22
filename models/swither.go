package models

import (
	"fmt"
	"net/smtp"

	"SenderS/env"
	"SenderS/models/messages"
	"github.com/streadway/amqp"
)

func ReaderSender(ch <-chan amqp.Delivery) error {
	auth := smtp.PlainAuth("", env.From, env.Pass, env.SmtpServ)

	for ms := range ch {
		// must be switch
		body, to, err := messages.TemplateOperation(ms)
		if err != nil {
			return err
		}
		err = smtp.SendMail(env.SmtpServ+":"+env.SmtpPort, auth, env.From, to, body.Bytes())
		if err != nil {
			return err
		}
		fmt.Println("Message send!")

	}
	return nil
}
