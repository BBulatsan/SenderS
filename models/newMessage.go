package models

import (
	"time"
)

func NewMessage(user string, email string, operation string) Message {
	id := int(time.Since(time.Now()))
	return Message{
		Id:        id,
		User:      user,
		Email:     email,
		Operation: operation,
	}
}
