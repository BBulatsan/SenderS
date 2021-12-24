package models

import (
	"fmt"
	"net/smtp"

	"SenderS/env"
	"SenderS/models/messages"
	"SenderS/modules/bus"
	"SenderS/modules/catcher"
)

func SwitcherSender(ch <-chan messages.Message) {
	auth := smtp.PlainAuth("", env.From, env.Pass, env.SmtpServ)
	addr := env.SmtpServ + ":" + env.SmtpPort
	for ms := range ch {
		switch ms.Rk {
		case bus.Operation:
			body, to, err := messages.TemplateOperation(ms.Body)
			if err != nil {
				catcher.HandlerError(err)
			}
			err = smtp.SendMail(addr, auth, env.From, to, body.Bytes())
			if err != nil {
				catcher.HandlerError(err)
			}
			fmt.Println("Message send!")
		case bus.Sale:
			body, to, err := messages.TemplateSale(ms.Body)
			if err != nil {
				catcher.HandlerError(err)
			}
			err = smtp.SendMail(addr, auth, env.From, to, body.Bytes())
			if err != nil {
				catcher.HandlerError(err)
			}
			fmt.Println("Message send!")
		default:
			fmt.Println("Unknown rk")
		}
	}
}
