package model

import (
	"fmt"
	"math/rand"
	"shareTravel/form"
	"time"
)

type Event form.Event

func (event *Event) CreateEvent() string {
	OpenSQL()
	//認証キーの取得
	key := createAuthKey()
	statement := "insert into event (auth_key,name,date,create_time) values(?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare error")
		return ""
	}
	t := time.Now()

	defer stmt.Close()
	stmt.Exec(key, event.Name, event.Date, t)

	if err != nil {
		fmt.Println("Exec error")

		return ""
	}
	return key
}

func GetEvent(event *Event) *Event {
	OpenSQL()

	err := Db.QueryRow("SELECT id,name,date FROM event WHERE auth_key = ?", event.AuthKey).Scan(&event.Id, &event.Name, &event.Date)

	if err != nil {
		fmt.Println("Exec error")
		panic(err.Error())
	}

	return event
}

func createAuthKey() string {
	var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	key := make([]rune, 16)
	for i := range key {
		key[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(key)
}
