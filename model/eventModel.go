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

	var err error

	if event.Id != 0 {
		err = Db.QueryRow("SELECT id,auth_key,name,date FROM event WHERE id = ?", event.Id).Scan(&event.Id, &event.AuthKey, &event.Name, &event.Date)
	} else if event.AuthKey != "" {
		err = Db.QueryRow("SELECT id,auth_key,name,date FROM event WHERE auth_key = ?", event.AuthKey).Scan(&event.Id, &event.AuthKey, &event.Name, &event.Date)
	} else {
		fmt.Println("it has no id and authkey")
		return nil
	}

	if err != nil {
		return nil
	}

	return event
}

func (event *Event) UpdatePool() {

	var pool int

	OpenSQL()
	//現在の端数プールを取得
	err := Db.QueryRow("SELECT pool FROM event WHERE id = ?", event.Id).Scan(&pool)

	if err != nil {
		//変数書き換えに失敗した場合終了
		fmt.Println("Exec error in SELECT")
		return
	}

	//レコードの更新処理
	statement := "UPDATE event SET pool = ?,update_time = ? WHERE id = ?;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare error")
		return
	}

	//現在時刻を取得
	t := time.Now()

	//処理終了後データソースを閉じる
	defer stmt.Close()

	//SQL実行
	stmt.Exec(pool, t, event.Id)

	if err != nil {
		fmt.Println("Exec error in UPDATE")
		return
	}
}

func createAuthKey() string {
	var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	key := make([]rune, 16)
	for i := range key {
		key[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(key)
}
