package model

import (
	"fmt"
	"shareTravel/form"
	"time"
)

type Member form.Member

func (member *Member) SaveMember() {
	OpenSQL()
	statement := "insert into member (event_id,name,create_time) values(?,?,?)"

	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare error")
	}
	t := time.Now()

	defer stmt.Close()
	stmt.Exec(member.EventId, member.Name, t)

	if err != nil {
		fmt.Println("Exec error")
	}
}
