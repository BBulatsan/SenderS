package models

import (
	"fmt"
	"net/smtp"

	"SenderS/env"
	"SenderS/models/messages"
	"SenderS/modules/bus"
	"SenderS/modules/catcher"
	"github.com/streadway/amqp"
)

func ReaderSender(ch <-chan amqp.Delivery) error {
	auth := smtp.PlainAuth("", env.From, env.Pass, env.SmtpServ)
	addr := env.SmtpServ + ":" + env.SmtpPort

	for ms := range ch {
		switch ms.RoutingKey {
		case bus.Operation:
			body, to, err := messages.TemplateOperation(ms)
			if err != nil {
				return err
			}
			err = smtp.SendMail(addr, auth, env.From, to, body.Bytes())
			if err != nil {
				return err
			}
			fmt.Println("Message send!")
		case bus.Sale:
			body, to, err := messages.TemplateSale(ms)
			if err != nil {
				return err
			}
			err = smtp.SendMail(addr, auth, env.From, to, body.Bytes())
			if err != nil {
				return err
			}
			fmt.Println("Message send!")
		default:
			fmt.Println("Unknown rk")
			err := ms.Nack(false, false)
			catcher.HandlerError(err)
		}
	}
	return nil
}
