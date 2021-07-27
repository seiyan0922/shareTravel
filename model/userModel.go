package model

import (
	"fmt"
	"shareTravel/form"
	"time"
)

type User form.User

var users []User

func (user *User) CreateUser() error {
	OpenSQL()
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

func IndexUser() []User {
	OpenSQL()

	rows, err := Db.Query("SELECT id,name,age,address FROM users")
	if err != nil {
		fmt.Println("Exec error")
		panic(err.Error())
	}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address)
		if err != nil {
			fmt.Println("Exec error")
			panic(err.Error())
		}
		users = append(users, user)
	}

	return users

}
