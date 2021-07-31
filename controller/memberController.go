package controller

import (
	"net/http"
	"shareTravel/form"
	"shareTravel/model"
	"strconv"
	"strings"
)

func MemberHandler(w http.ResponseWriter, r *http.Request, path string) {
	arr := strings.Split(path, "/")
	switch arr[0] {
	case "add":
		memberAddHandler(w, r)
	case "save":
		memberSaveHandler(w, r)
	}
}

func memberAddHandler(w http.ResponseWriter, r *http.Request) {
	event := new(form.Event)
	query := r.URL.Query().Encode()
	strid := strings.Split(query, "=")[1]
	event.Id, _ = strconv.Atoi(strid)

	RenderTemplate(w, "view/member/add", event)
}

func memberSaveHandler(w http.ResponseWriter, r *http.Request) {
	member := new(model.Member)
	member.Name = r.FormValue("name")

	query := r.URL.Query().Encode()
	strid := strings.Split(query, "=")[1]
	member.EventId, _ = strconv.Atoi(strid)

	member.SaveMember()

	RenderTemplate(w, "view/member/complete", member)
}
