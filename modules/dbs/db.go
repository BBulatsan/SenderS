package dbs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/smtp"
	"time"

	"SenderS/env"
	"SenderS/models/messages"
	"SenderS/modules/bus"
	"SenderS/modules/publisher"
	_ "github.com/mattn/go-sqlite3"
)

type DbConn struct {
	conn *sql.DB
}

type MessageDb struct {
	Rk   string
	Mess []byte
}

type item struct {
	id   int
	rk   string
	body string
	read int
}

//CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY, 'RK' message_text, Body message_text,Type message_text, Read TINYINT(1));

func (d *DbConn) InitDb() {
	db, err := sql.Open("sqlite3", "./modules/dbs/dbs/item.db")
	if err != nil {
		panic(err)
	}
	d.conn = db
}

func (d *DbConn) AddMessage(ms publisher.Message) error {
	body, err := json.Marshal(ms.Mess)
	if err != nil {
		return err
	}
	statement, err := d.conn.Prepare("INSERT INTO messages (RK, Body,Read) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(ms.Rk, body, 0)
	if err != nil {
		return err
	}
	return nil
}

func (d *DbConn) CheckTakeMessages() error {
	auth := smtp.PlainAuth("", env.From, env.Pass, env.SmtpServ)
	addr := env.SmtpServ + ":" + env.SmtpPort
	for {
		rows, err := d.conn.Query("SELECT *  FROM messages WHERE Read = 0")
		if err != nil {
			return err
		}
		items := []item{}
		for rows.Next() {
			i := item{}
			err := rows.Scan(&i.id, &i.rk, &i.body, &i.read)
			if err != nil {
				fmt.Println(err)
				continue
			}
			items = append(items, i)
		}
		for _, i := range items {
			switch i.rk {
			case bus.Operation:
				body, to, err := messages.TemplateOperation([]byte(i.body))
				if err != nil {
					return err
				}
				err = smtp.SendMail(addr, auth, env.From, to, body.Bytes())
				if err != nil {
					return err
				}
				fmt.Println("Message send!")
			case bus.Sale:
				body, to, err := messages.TemplateSale([]byte(i.body))
				if err != nil {
					return err
				}
				err = smtp.SendMail(addr, auth, env.From, to, body.Bytes())
				if err != nil {
					return err
				}
				fmt.Println("Message send!")
			default:
				fmt.Println("Unknown rk")
			}
			_, err := d.conn.Exec("update messages set Read = 1 where id = $1", i.id)
			if err != nil {
				return err
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
