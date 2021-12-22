package publisher

import (
	"encoding/json"

	"SenderS/env"
	"github.com/streadway/amqp"
)

func Publisher(ch *amqp.Channel, ms interface{}) error {
	body, err := json.Marshal(ms)
	if err != nil {
		return err
	}
	err = ch.Publish(env.Exchange, "*.email.#", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		return err
	}
	return nil
}
