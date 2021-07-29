package controller

import (
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request, path string) {
	// p, err := LoadPage(title)

	// user := model.User{Name: "Test", Age: 20, Address: 1999999}

	// if sqlerr := user.CreateUser(); sqlerr != nil {
	// 	fmt.Println(sqlerr)
	// }

	// if err != nil {
	// 	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	// 	return
	// }
	// RenderTemplate(w, "view", p)

}
