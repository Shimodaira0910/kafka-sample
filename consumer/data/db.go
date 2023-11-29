package data

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type Data struct{
	db *sql.DB
}

const (
	conn = "host=localhost port=5432 user=test password=test dbname=test sslmode=disable options='--client_encoding=UTF8'"
)

func NewData() *Data {
   db, err := sql.Open("postgres", conn)
   if err != nil {
       log.Fatalln("connection to database failed!!:", err)
   }
   return &Data{db: db}
}

func (d *Data) CloseDb() {
	if d.db != nil{
		d.db.Close()
	} 
}

func (d *Data) InsertQuery(consumed int, msg string) error {
	if d.db == nil {
		return fmt.Errorf("データベース接続が初期化されていません")
	}
	replaceMessage := d.replaceNullByte(msg)
	_, err := d.db.Exec("INSERT INTO test (id, value) VALUES ($1, $2)", consumed, replaceMessage)
	return err
}

func (d *Data) replaceNullByte(msg string) string{
    replacedMessage := strings.ReplaceAll(msg, "\x00", "")
    return replacedMessage
}
