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
	http.HandleFunc("/save/", controller.MakeHandler(controller.SaveHandler))
	http.HandleFunc("/create/", controller.MakeHandler(controller.CreateUserHandler))
	http.HandleFunc("/index/", controller.MakeHandler(controller.IndexUserHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
