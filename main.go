package main

import (
	"fmt"

	"SenderS/models"
	"SenderS/modules/bus"
	"SenderS/modules/publisher"
)

func main() {
	conn, queue := bus.InitRebbit()

	message := models.NewMessage("Bohdan", "bogdan315991@gmail.com", "registration")
	publisher.Publisher(conn, message)

	MessagesChan := bus.Consume(queue, conn)
	go models.ReaderSender(MessagesChan)

	var ex string
	fmt.Scan(&ex)
	err := conn.Close()
	if err != nil {
		panic(err)
	}
}
