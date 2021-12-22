package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"time"

	"github.com/streadway/amqp"
)

type MessageOperation struct {
	Id        int
	User      string
	Email     string
	Operation string
}

func TemplateOperation(ms amqp.Delivery) (bytes.Buffer, []string, error) {
	var dat MessageOperation
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
	tmpl, err := template.ParseFiles("template/reg.html")
	if err != nil {
		return bytes.Buffer{}, nil, nil
	}

	body.Write([]byte(fmt.Sprintf("Subject: This is text about %s  \n%s\n\n", dat.Operation, mimeHeaders)))
	err = tmpl.Execute(&body, dat)
	if err != nil {
		return bytes.Buffer{}, nil, nil
	}

	to := []string{dat.Email}

	return body, to, nil
}

func NewMessageOperation(user string, email string, operation string) MessageOperation {
	id := int(time.Since(time.Now()))
	return MessageOperation{
		Id:        id,
		User:      user,
		Email:     email,
		Operation: operation,
	}
}
