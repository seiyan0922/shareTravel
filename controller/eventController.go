package controller

import (
	"fmt"
	"net/http"
	"shareTravel/form"
	"shareTravel/model"
	"strings"
)

func EventHandler(w http.ResponseWriter, r *http.Request, path string) {
	arr := strings.Split(path, "/")
	switch arr[0] {
	case "create":
		createEventHandler(w, r)
	case "confirm":
		confirmEventHandler(w, r)
	case "save":
		saveEventHandler(w, r)

	}
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "view/event/create", nil)
}

func confirmEventHandler(w http.ResponseWriter, r *http.Request) {
	event := new(form.Event)

	event.Name = r.FormValue("name")
	event.Date = r.FormValue("datetime")

	RenderTemplate(w, "view/event/confirm", event)

}

func saveEventHandler(w http.ResponseWriter, r *http.Request) {

	event := new(model.Event)
	event.Name = r.FormValue("name")
	event.Date = r.FormValue("date")
	fmt.Println(r.FormValue("name"))
	event.CreateEvent()
	RenderTemplate(w, "view/event/complete", event)

}
