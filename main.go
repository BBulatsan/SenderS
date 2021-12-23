package main

import (
	"fmt"

	"SenderS/models"
	"SenderS/models/messages"
	"SenderS/modules/bus"
	"SenderS/modules/catcher"
	"SenderS/modules/publisher"
)

func main() {
	RConn := bus.RConn{}
	RConn.InitRabbit()

	//messageSale := messages.NewMessageSale("Bohdan", "bogdan315991@gmail.com", 20)
	//message := publisher.Message{Rk: "*.test.#", Mess: messageSale}
	messageOperation := messages.NewMessageOperation("Bohdan", "bogdan315991@gmail.com", "registration")
	message := publisher.Message{Rk: bus.Operation, Mess: messageOperation}
	message2 := publisher.Message{Rk: bus.Operation, Mess: messageOperation}
	//messageSale := messages.NewMessageSale("Bohdan", "bogdan315991@gmail.com", 25)
	//message := publisher.Message{Rk: bus.Sale, Mess: messageSale}

	err := message.Publisher(RConn.Conn)
	catcher.HandlerError(err)
	err = message2.Publisher(RConn.Conn)
	catcher.HandlerError(err)

	err = RConn.Consume()
	catcher.HandlerError(err)

	go models.ReaderSender(RConn.MessagesChan)

	var ex string
	fmt.Scan(&ex)

}
