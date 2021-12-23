package main

import (
	"fmt"

	"SenderS/models"
	"SenderS/modules/bus"
	"SenderS/modules/dbs"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	RConn := bus.RConn{}
	RConn.InitRabbit()
	DConn := dbs.DbConn{}
	DConn.InitDb()
	go models.ReaderSenderRabbit(RConn.MessagesChan)
	go DConn.CheckTakeMessages()

	//messageSale := messages.NewMessageSale("Bohdan", "bogdan315991@gmail.com", 20)
	//message := publisher.Message{Rk: "*.test.#", Mess: messageSale}
	//messageOperation := messages.NewMessageOperation("Bohdan", "bogdan315991@gmail.com", "registration")
	//message := publisher.Message{Rk: bus.Operation, Mess: messageOperation}
	//messageSale := messages.NewMessageSale("Bohdan", "bogdan315991@gmail.com", 25)
	//message := publisher.Message{Rk: bus.Sale, Mess: messageSale}

	//err := message.Publisher(RConn.Conn)
	//catcher.HandlerError(err)
	//err := DConn.AddMessage(message)
	//catcher.HandlerError(err)

	var ex string
	fmt.Scan(&ex)

}
