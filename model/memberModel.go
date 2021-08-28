package model

import (
	"fmt"
	"time"
)

type Member struct {
	Id          int
	EventId     int
	Name        string
	Temporarily int
	Calculate   int
	Total       int
}

func (member *Member) SaveMember() {

	statement := "insert into member (event_id,name,create_time) values(?,?,?)"

	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
	}
	t := time.Now()

	defer stmt.Close()
	stmt.Exec(member.EventId, member.Name, t)

	if err != nil {
		fmt.Println(err)
	}
}

func GetMembers(id int) []Member {

	rows, err := Db.Query("SELECT id,name from member WHERE event_id = ?", id)

	if err != nil {
		fmt.Println(err)
	}

	var members []Member

	for rows.Next() {
		member := Member{}
		err := rows.Scan(&member.Id, &member.Name)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		members = append(members, member)
	}

	return members
}

func (member *Member) GetMemberTemporarily() {

	rows, err := Db.Query("SELECT total from expense WHERE temporarily_member = ?", member.Id)

	if err != nil {
		fmt.Println(err)
	}

	var total_temporarily int

	for rows.Next() {
		var temporarily int
		err := rows.Scan(&temporarily)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		total_temporarily += temporarily
	}

	member.Temporarily = total_temporarily
}
