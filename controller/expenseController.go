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

func ExpenseHandler(w http.ResponseWriter, r *http.Request, path string) {
	arr := strings.Split(path, "/")
	switch arr[0] {
	case "add":
		addExpenseHandler(w, r)
	case "complete":
		completeExpenseHandler(w, r)
	case "save":
		saveExpenseHandler(w, r)
	case "show":
		showExpenseHandler(w, r)
	}
}

func addExpenseHandler(w http.ResponseWriter, r *http.Request) {

	event_id := common.GetQueryParam(r)
	event := new(form.Event)
	event.Id, _ = strconv.Atoi(event_id)
	RenderTemplate(w, "view/expense/add", event)

}

func completeExpenseHandler(w http.ResponseWriter, r *http.Request) {

	//POSTデータ：金額を取得
	total, _ := strconv.Atoi(r.FormValue("price"))

	//POSTデータ：名前を取得
	name := r.FormValue("name")

	//POSTデータ：備考を取得
	remarks := r.FormValue("remarks")

	//支払いポインタ
	expense := new(model.Expense)

	//データの格納
	event_id, _ := strconv.Atoi(common.GetQueryParam(r))

	expense.EventId = event_id

	expense.Name = name

	expense.Total = total

	expense.Remarks = remarks

	//支払い情報を保存
	expense.Id = expense.AddExpense()

	fmt.Println(expense.Id)

	//各参加者の負担金の登録
	//イベントIDから参加者データを取得
	members := model.GetMembers(event_id)
	//参加者人数を取得
	member_count := len(members)

	//各参加者の負担金を設定
	each_price := total / member_count

	//金額が参加人数で割り切れない場合
	if fraction := total % member_count; fraction != 0 {
		event := new(model.Event)
		event.Id = event_id
		event.UpdatePool()
	}

	//イベント参加者各位の負担金額を保存する
	err := model.CreateMemberExpense(event_id, expense, each_price)
	if err != nil {
		fmt.Println(err)
	}

	RenderTemplate(w, "view/expense/complete", nil)

}

func saveExpenseHandler(w http.ResponseWriter, r *http.Request) {

}

func showExpenseHandler(w http.ResponseWriter, r *http.Request) {

}
