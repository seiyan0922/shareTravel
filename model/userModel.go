package model

import (
	"fmt"
	"shareTravel/form"
	"time"
)

type User form.User

func (user *User) CreateUser() error {
	OpenSQL()
	fmt.Println(user)
	statement := "insert into users (name,age,address,create_time) values(?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare error")
		return err
	}
	t := time.Now()

	defer stmt.Close()
	stmt.Exec(user.Name, user.Age, user.Address, t)

	if err != nil {
		fmt.Println("Exec error")

		return err
	}
	return err
}
