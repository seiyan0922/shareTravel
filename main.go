package main

import (
	"log"
	"net/http"

	"shareTravel/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/view/", controller.MakeHandler(controller.ViewHandler))
	http.HandleFunc("/edit/", controller.MakeHandler(controller.EditHandler))
	http.HandleFunc("/save/", controller.MakeHandler(controller.SaveHandler))
	http.HandleFunc("/create/", controller.MakeHandler(controller.CrateUserHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
