package main

import (
	"fmt"

	"SenderS/env"
	"SenderS/models"
	"SenderS/models/messages"
	"SenderS/modules/bus"
	"SenderS/modules/catcher"
	"SenderS/modules/dbs"
	"SenderS/modules/publisher"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	RConn := bus.RConn{}
	DConn := dbs.DbConn{}
	MessageCh := make(chan messages.Message)
	go models.SwitcherSender(MessageCh)

	switch env.Mode {
	case "Rabbit":
		RConn.InitRabbit()
		go bus.ReaderRabbit(RConn.MessagesChan, MessageCh)
	case "DB":
		DConn.InitDb()
		go DConn.CheckTakeMessages(MessageCh)
	case "ALL":
		RConn.InitRabbit()
		DConn.InitDb()
		go bus.ReaderRabbit(RConn.MessagesChan, MessageCh)
		go DConn.CheckTakeMessages(MessageCh)
	default:
		fmt.Println("no config! env.Mode is empty")
	}

	messageSale := messages.NewMessageSale("Bohdan", "bogdan315991@gmail.com", 20)
	message := publisher.Message{Rk: "*.test.#", Mess: messageSale}
	//messageOperation := messages.NewMessageOperation("Bohdan", "bogdan315991@gmail.com", "registration")
	//message := publisher.Message{Rk: bus.Operation, Mess: messageOperation}
	//messageSale := messages.NewMessageSale("Bohdan", "bogdan315991@gmail.com", 25)
	message2 := publisher.Message{Rk: bus.Sale, Mess: messageSale}

	err := message.Publisher(RConn.Conn)
	catcher.HandlerError(err)
	err = message2.Publisher(RConn.Conn)
	catcher.HandlerError(err)
	//err = DConn.AddMessage(message)
	//catcher.HandlerError(err)

	var ex string
	fmt.Scan(&ex)

}
