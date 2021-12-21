package publisher

import (
	"encoding/json"

	"SenderS/env"
	"SenderS/models"
	"github.com/streadway/amqp"
)

func Publisher(ch *amqp.Channel, ms models.Message) {
	body, err := json.Marshal(ms)
	if err != nil {
		panic(err)
	}
	err = ch.Publish(env.Exchange, "*.email.#", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		panic(err)
	}
}
