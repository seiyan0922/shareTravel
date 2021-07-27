package controller

import (
	"fmt"
	"net/http"
	"shareTravel/model"
	"strconv"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, title string) {

	//HTTPメソッドを取得
	method := r.Method

	//モデルよりUser構造を取得(大元はFormPackage)
	user := new(model.User)

	//HTTPメソッドにより処理を分岐
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

		user.CreateUser()
		// RenderTemplate(w, "complete", user)
		RenderTemplate(w, "complete", user)

	}

}

func IndexUserHandler(w http.ResponseWriter, r *http.Request, title string) {
	users := model.IndexUser()

	RenderTemplate(w, "index", users)

}
