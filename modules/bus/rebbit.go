package bus

import (
	"SenderS/env"
	"github.com/streadway/amqp"
)

func InitRebbit() (*amqp.Channel, amqp.Queue) {
	conn := Conn("amqp://rabbitmq:rabbitmq@localhost:5672/")
	queue := AddQueue(conn, env.Queue)
	AddExchange(conn, env.Exchange)
	AddBind(conn, queue, "*.email.#", env.Exchange)
	return conn, queue
}

func Conn(url string) *amqp.Channel {
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	amqpChannel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return amqpChannel
}

func AddQueue(channel *amqp.Channel, name string) amqp.Queue {
	queue, err := channel.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	return queue
}

func AddExchange(channel *amqp.Channel, name string) {
	err := channel.ExchangeDeclare(name, "topic", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}

func AddBind(channel *amqp.Channel, queue amqp.Queue, rk string, exchange string) {
	err := channel.QueueBind(queue.Name, rk, exchange, false, nil)
	if err != nil {
		panic(err)
	}
}

func Consume(queue amqp.Queue, channel *amqp.Channel) <-chan amqp.Delivery {
	messageChannel, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	return messageChannel
}
