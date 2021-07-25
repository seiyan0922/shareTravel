package controller

import (
	"fmt"
	"net/http"
	"shareTravel/model"
	"strconv"
	"time"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request, title string) {

	method := r.Method

	user := new(User)

	switch method {
	case "GET":
		RenderTemplate(w, "create", user)
	case "POST":
		user.Name = r.FormValue("name")
		age, err := strconv.Atoi(r.FormValue("age"))

		//文字列変換時のエラー処理
		if err != nil {
			fmt.Println("error: conv int userAge")
			return
		}
		user.Age = age
		user.Address = r.FormValue("address")
		// user := user{Name: ""}

		CreateUser(user)
		// RenderTemplate(w, "complete", user)
		RenderTemplate(w, "complete", user)

	}

}

func CreateUser(user *User) error {
	model.OpenSQL()
	fmt.Println("CreateUser通ってる")
	fmt.Println(user)
	statement := "insert into users (name,age,address,create_time) values(?,?,?,?)"
	stmt, err := model.Db.Prepare(statement)
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
