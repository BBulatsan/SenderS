package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"time"

	"github.com/streadway/amqp"
)

type MessageSale struct {
	Id       int
	User     string
	Email    string
	Percent  uint8
	PromoCod string
}

func TemplateSale(ms amqp.Delivery) (bytes.Buffer, []string, error) {
	var dat MessageSale
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	err := json.Unmarshal(ms.Body, &dat)
	if err != nil {
		return bytes.Buffer{}, nil, nil
	}
	err = ms.Ack(true)
	if err != nil {
		return bytes.Buffer{}, nil, nil
	}
	tmpl, err := template.ParseFiles("template/sale.html")
	if err != nil {
		return bytes.Buffer{}, nil, nil
	}

	body.Write([]byte(fmt.Sprintf("Subject: You have a new Sale!  \n%s\n\n", mimeHeaders)))
	err = tmpl.Execute(&body, dat)
	if err != nil {
		return bytes.Buffer{}, nil, nil
	}

	to := []string{dat.Email}

	return body, to, nil
}

func NewMessageSale(user string, email string, percent uint8) MessageSale {
	id := int(time.Since(time.Now()))
	cod := "Goods"
	return MessageSale{
		Id:       id,
		User:     user,
		Email:    email,
		Percent:  percent,
		PromoCod: cod,
	}
}
