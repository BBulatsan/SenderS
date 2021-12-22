package bus

import (
	"SenderS/env"
	"SenderS/modules/catcher"
	"github.com/streadway/amqp"
)

type RConn struct {
	Conn         *amqp.Channel
	Queue        amqp.Queue
	MessagesChan <-chan amqp.Delivery
}

func (r *RConn) InitRabbit() {
	err := r.conn(env.HostRebbit)
	catcher.HandlerError(err)
	err = r.addQueue(env.Queue)
	catcher.HandlerError(err)
	err = r.addExchange(env.Exchange)
	catcher.HandlerError(err)
	for _, b := range Bindings() {
		err = r.addBind(b, env.Exchange)
		catcher.HandlerError(err)
	}
}

func (r *RConn) conn(url string) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	amqpChannel, err := conn.Channel()
	if err != nil {
		return err
	}
	r.Conn = amqpChannel
	return nil
}

func (r *RConn) addQueue(name string) error {
	queue, err := r.Conn.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		return err
	}
	r.Queue = queue
	return nil
}

func (r *RConn) addExchange(name string) error {
	err := r.Conn.ExchangeDeclare(name, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *RConn) addBind(rk string, exchange string) error {
	err := r.Conn.QueueBind(r.Queue.Name, rk, exchange, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *RConn) Consume() error {
	messageChannel, err := r.Conn.Consume(
		r.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	r.MessagesChan = messageChannel
	return nil
}
