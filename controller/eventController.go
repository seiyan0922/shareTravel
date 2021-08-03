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
	case "show":
		showEventHandler(w, r)

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
	key := event.CreateEvent()
	event.AuthKey = key
	RenderTemplate(w, "view/event/complete", event)

}

func showEventHandler(w http.ResponseWriter, r *http.Request) {

	event := new(model.Event)
	query := r.URL.Query()
	param := query.Encode()
	event.AuthKey = strings.Split(param, "=")[1]
	event = model.GetEvent(event)
	members := model.GetMembers(event.Id)

	status := make(map[string]interface{})
	status["Event"] = event
	status["Members"] = members

	RenderTemplate(w, "view/event/show", status)

}
