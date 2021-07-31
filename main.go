package main

import (
	"log"
	"net/http"
	"shareTravel/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/top/", controller.MakeHandler(controller.TopHandler))
	http.HandleFunc("/view/", controller.MakeHandler(controller.ViewHandler))
	http.HandleFunc("/edit/", controller.MakeHandler(controller.EditHandler))
	http.HandleFunc("/create/", controller.MakeHandler(controller.CreateUserHandler))
	http.HandleFunc("/index/", controller.MakeHandler(controller.IndexUserHandler))
	http.HandleFunc("/event/", controller.MakeHandler(controller.EventHandler))
	http.HandleFunc("/member/", controller.MakeHandler(controller.MemberHandler))

	//cssファイルへのハンドラを定義
	http.Handle("/layout/", http.StripPrefix("/layout/", http.FileServer(http.Dir("layout/"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
