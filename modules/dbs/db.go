package dbs

import (
	"database/sql"
	"encoding/json"
	"time"

	"SenderS/models/messages"
	"SenderS/modules/catcher"
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
	db, err := sql.Open("sqlite3", "./modules/dbs/dbstore/item.db")
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
	statement, err := d.conn.Prepare("INSERT INTO messages (rk, body, read) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(ms.Rk, body, 0)
	if err != nil {
		return err
	}
	return nil
}

func (d *DbConn) CheckTakeMessages(chIn chan messages.Message) {
	for {
		rows, err := d.conn.Query("SELECT *  FROM messages WHERE read = 0")
		if err != nil {
			catcher.HandlerError(err)
		}
		items := []item{}
		for rows.Next() {
			i := item{}
			err := rows.Scan(&i.id, &i.rk, &i.body, &i.read)
			if err != nil {
				catcher.HandlerError(err)
				continue
			}
			items = append(items, i)
		}
		for _, i := range items {
			chIn <- messages.Message{Rk: i.rk, Body: []byte(i.body)}
			_, err := d.conn.Exec("update messages set read = 1 where id = $1", i.id)
			if err != nil {
				catcher.HandlerError(err)
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
