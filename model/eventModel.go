package model

import (
	"fmt"
	"shareTravel/form"
	"time"
)

type Event form.Event

func (event *Event) CreateEvent() error {
	OpenSQL()
	statement := "insert into event (name,date,create_time) values(?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare error")
		return err
	}
	t := time.Now()

	defer stmt.Close()
	stmt.Exec(event.Name, event.Date, t)

	if err != nil {
		fmt.Println("Exec error")

		return err
	}
	return err
}
