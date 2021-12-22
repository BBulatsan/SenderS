package main

import (
	"fmt"

	"SenderS/models"
	"SenderS/models/messages"
	"SenderS/modules/bus"
	"SenderS/modules/cather"
	"SenderS/modules/publisher"
)

func main() {
	RConn := bus.RConn{}
	RConn.InitRabbit()

	message := messages.NewMessageOperation("Bohdan", "bogdan315991@gmail.com", "registration")
	err := publisher.Publisher(RConn.Conn, message)
	cather.HandlerError(err)
	err = RConn.Consume()
	cather.HandlerError(err)
	go models.ReaderSender(RConn.MessagesChan)

	var ex string
	fmt.Scan(&ex)

}
