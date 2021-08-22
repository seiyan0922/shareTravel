package controller

import (
	"fmt"
	"net/http"
	"shareTravel/common"
	"shareTravel/form"
	"shareTravel/model"
	"strconv"
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
	case "indexMember":
		showMembersEventHandler(w, r)
	case "edit":
		editEventHandler(w, r)
	case "download":
		csvDownLoad(w, r)
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
	query := r.URL.Query().Encode()
	qarr := strings.Split(query, "=")

	if qarr[0] == "event_id" {
		event.Id, _ = strconv.Atoi(qarr[1])
	} else {
		event.AuthKey = qarr[1]
	}

	event = model.GetEvent(event)

	showEventRender(w, event)

}

//イベントTOP表示共通処理
func showEventRender(w http.ResponseWriter, event *model.Event) {

	status := make(map[string]interface{})
	status["Event"] = &event

	expenses := event.GetExpensesByEventId()
	if len(expenses) == 0 {
		status["Expenses"] = nil

	} else {
		status["Expenses"] = &expenses
	}

	RenderTemplate(w, "view/event/show", status)

}

func showMembersEventHandler(w http.ResponseWriter, r *http.Request) {
	event := new(model.Event)
	event_id, _ := strconv.Atoi(common.GetQueryParam(r))
	event.Id = event_id

	event = model.GetEvent(event)

	members := model.GetMembers(event_id)

	//各参加者の立替金を取得
	for i := 0; i < len(members); i++ {
		members[i].GetMemberTemporarily()
	}

	nomember_flg := false
	if len(members) == 0 {
		nomember_flg = true
	} else {
		//参加メンバー負担金総額取得処理
		GetMembersTotal(members)
	}
	showMembersEventRender(w, event, &members, nomember_flg)

}

func showMembersEventRender(w http.ResponseWriter, event *model.Event, members *[]model.Member, nomember_flg bool) {

	status := make(map[string]interface{})
	status["Event"] = event

	if !nomember_flg {
		status["Members"] = members
	} else {
		status["Members"] = nil
	}

	RenderTemplate(w, "view/event/showMember", status)

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

//イベント設計ページ
func editEventHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Encode()
	qarr := strings.Split(query, "=")
	event := new(model.Event)
	event.Id, _ = strconv.Atoi(qarr[1])

	//リクエストメソッドによる条件分岐
	switch r.Method {
	case "GET":
		event = model.GetEvent(event)

		//イベント編集テンプレートの読み込み
		RenderTemplate(w, "view/event/edit", event)

	case "POST":
		event.AuthKey = r.FormValue("auth_key")
		event.Name = r.FormValue("name")
		event.Date = r.FormValue("date")
		fmt.Println(r.FormValue("date"))
		event.UpdateEvent()

		showEventRender(w, event)

	}
}

func csvDownLoad(w http.ResponseWriter, r *http.Request) {

	event := new(model.Event)

	event.Id, _ = strconv.Atoi(common.GetQueryParam(r))

	event = model.GetEvent(event)

	members := model.GetMembers(event.Id)

	expenses := event.GetExpensesByEventId()

	head_line := event.Name + ",\n" + "端数合計" + strconv.Itoa(event.Pool) + "\n"

	expense_lines := ","

	total_expense_line := "合計（端数込み）,"

	for _, expense := range expenses {
		expense_lines += expense.Name + ","
		total_expense_line += strconv.Itoa(expense.Total) + "円,"

	}
	expense_lines += "個人負担合計,立替,請求合計\n"

	var member_lines string

	for _, member := range members {
		member_lines += member.Name + ","
		temp := 0
		for _, expense := range expenses {
			member.SearchMemberExpense(expense.Id)
			member_lines += strconv.Itoa(member.Calculate) + "円,"
			if expense.TemporarilyMemberId == member.Id {
				temp += expense.Total
			}
		}
		model.GetMemberExpense(&member)
		member_lines += strconv.Itoa(member.Total) + "円," + strconv.Itoa(temp) + "円," +
			strconv.Itoa(member.Total-temp) + "円\n"
	}

	csv_string := head_line + expense_lines + member_lines + total_expense_line
	out := []byte(csv_string)

	// ファイル名
	w.Header().Set("Content-Disposition", "attachment; filename=result.csv")
	// コンテントタイプ
	w.Header().Set("Content-Type", "text/csv")
	// ファイルの長さ
	w.Header().Set("Content-Length", string(len(out)))
	// bodyに書き込み
	w.Write(out)
}
