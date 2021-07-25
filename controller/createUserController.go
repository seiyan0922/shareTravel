package controller

import (
	"net/http"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request, title string) {

	method := r.Method

	user := &User{
		Name:    "山田",
		Age:     12,
		Address: "sadf",
	}

	switch method {
	case "GET":
		RenderTemplate(w, "create", user)
	case "POST":

	}

}
