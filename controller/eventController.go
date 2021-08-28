package controller

import (
	"net/http"
	"shareTravel/common"
	"shareTravel/model"
	"shareTravel/validate"
	"strconv"
	"strings"
)

func EventHandler(w http.ResponseWriter, r *http.Request, path string) {

	//pathによって実行する関数を分岐
	arr := strings.Split(path, common.SLASH)
	switch arr[common.ZERO] {
	case CREATE:
		createEventHandler(w, r)
	case CONFIRM:
		confirmEventHandler(w, r)
	case SAVE:
		saveEventHandler(w, r)
	case SHOW:
		showEventHandler(w, r)
	case SEARCH:
		searchEventHandler(w, r)
	case INDEX_MEMBER:
		showMembersEventHandler(w, r)
	case EDIT:
		editEventHandler(w, r)
	case DOWNLOAD:
		csvDownLoad(w, r)
	}
}

//新規イベント作成画面表示
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, EVENT_CREATE_PATH, nil)
}

//新規イベント確認画面表示
func confirmEventHandler(w http.ResponseWriter, r *http.Request) {

	//入力値を構造体に変換
	event, _ := formValueEncodeForEvent(r)

	//バリデーションチェック
	if ok, errs := validate.EventValidater(event); !ok {
		//バリデーションエラーがあった場合
		errorHandler(w, EVENT_CREATE_PATH, nil, errs)
	}

	//クライアントへの連携値を整形
	status := autoMapperForView(event)

	RenderTemplate(w, EVENT_CONFIRM_PATH, status)
}

//新規イベント作成完了画面表示
func saveEventHandler(w http.ResponseWriter, r *http.Request) {

	//入力値を構造体に変換
	event, _ := formValueEncodeForEvent(r)

	if ok, errs := validate.EventValidater(event); !ok {
		//バリデーションエラーがあった場合
		errorHandler(w, EVENT_CREATE_PATH, nil, errs)
	}

	//DB登録処理
	event = event.CreateEvent()

	//クライアントへの連携値を整形
	status := autoMapperForView(event)

	RenderTemplate(w, EVENT_COMPLETE_PATH, status)

}

//リクエストをもとに入力値をイベント用構造体に変換
func formValueEncodeForEvent(r *http.Request) (*model.Event, error) {
	event := new(model.Event)

	//POST値が存在するかの判別
	if name := r.FormValue("name"); name != EMPTY {
		event.Name = name
	}

	if pool := r.FormValue("pool"); pool != EMPTY {
		event.Pool, _ = strconv.Atoi(pool)
	}

	if datetime := r.FormValue("date"); datetime != EMPTY {
		event.Date = datetime
	}

	return event, nil

}

//イベントTOP表示ハンドラー
func showEventHandler(w http.ResponseWriter, r *http.Request) {

	//構造体ポインタの作成
	event := new(model.Event)

	//クエリパラメータの取得
	id_str := common.GetQueryParam(r)

	//クエリパラメータを整数型に変換
	event.Id, _ = strconv.Atoi(id_str)

	//イベントIDに紐づくイベントを取得
	event.GetEvent()

	//イベントTOP画面をレンダリング
	showEventRender(w, event)

}

//イベントTOP画面レンダー
func showEventRender(w http.ResponseWriter, event *model.Event) {

	expenses := event.GetExpensesByEventId()

	status := autoMapperForView(event, expenses)

	RenderTemplate(w, "view/event/show", status)

}

//
//
//
//
//
//
//
//リファクタリング未済
//
//
//
//
//
//
//

func showMembersEventHandler(w http.ResponseWriter, r *http.Request) {
	event := new(model.Event)
	event_id, _ := strconv.Atoi(common.GetQueryParam(r))
	event.Id = event_id

	event.GetEvent()

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
		event.GetEvent()

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
		event.GetEvent()

		//イベント編集テンプレートの読み込み
		RenderTemplate(w, "view/event/edit", event)

	case "POST":
		event.AuthKey = r.FormValue("auth_key")
		event.Name = r.FormValue("name")
		event.Date = r.FormValue("date")

		event.UpdateEvent()

		showEventRender(w, event)

	}
}

func csvDownLoad(w http.ResponseWriter, r *http.Request) {

	event := new(model.Event)

	event.Id, _ = strconv.Atoi(common.GetQueryParam(r))

	event.GetEvent()

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
