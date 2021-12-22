package publisher

import (
	"encoding/json"

	"SenderS/env"
	"github.com/streadway/amqp"
)

type Message struct {
	Rk   string
	Mess interface{}
}

func (m *Message) Publisher(ch *amqp.Channel) error {
	body, err := json.Marshal(m.Mess)
	if err != nil {
		return err
	}
	err = ch.Publish(env.Exchange, m.Rk, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		return err
	}
	return nil
}
