package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"time"
)

type MessageSale struct {
	Id       int
	User     string
	Email    string
	Percent  uint8
	PromoCod string
}

func TemplateSale(ms []byte) (bytes.Buffer, []string, error) {
	var dat MessageSale
	var body bytes.Buffer
	var tmpl *template.Template

	err := json.Unmarshal(ms, &dat)
	if err != nil {
		return bytes.Buffer{}, nil, nil
	}
	if cache.sale != nil {
		tmpl = cache.sale
	} else {
		tmpl, err := template.ParseFiles("template/sale.html")
		if err != nil {
			return bytes.Buffer{}, nil, nil
		}
		cache.sale = tmpl

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
