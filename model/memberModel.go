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

func GetMembers(id int) []Member {

	OpenSQL()

	rows, err := Db.Query("SELECT id,name from member WHERE event_id = ?", id)

	if err != nil {
		fmt.Println("Query Error")
	}

	var members []Member

	for rows.Next() {
		member := Member{}
		err := rows.Scan(&member.Id, &member.Name)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		members = append(members, member)
	}

	return members
}
