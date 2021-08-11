package controller

import (
	"fmt"
	"net/http"
	"shareTravel/common"
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
	case "search":
		searchEventHandler(w, r)

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
	event.AuthKey = common.GetQueryParam(r)
	event = model.GetEvent(event)

	showEventRender(w, event)

}

//イベントTOP表示共通処理
func showEventRender(w http.ResponseWriter, event *model.Event) {

	expenses := event.GetExpensesByEventId()
	status := make(map[string]interface{})
	status["Event"] = &event
	status["Expenses"] = &expenses
	fmt.Println(status["Expenses"])

	RenderTemplate(w, "view/event/show", status)

}

func searchEventHandler(w http.ResponseWriter, r *http.Request) {

	//リクエストメソッドによる条件分岐
	switch r.Method {
	case "GET":
		//GETの場合テンプレートを読み込み
		RenderTemplate(w, "view/event/search", nil)
	case "POST":
		//POSTの場合認証キーから該当のイベントを検索
		auth_key := r.FormValue("auth_key")

		event := new(model.Event)
		event.AuthKey = auth_key

		event = model.GetEvent(event)

		//イベント取得に成功した場合
		if event != nil {
			//イベントTOPページの読み込み
			showEventRender(w, event)

		} else {
			RenderTemplate(w, "view/event/search", nil)
		}

	}
}
