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
	fmt.Println("コントローラーきた")
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

	//支払い構造体ポインタ
	expense := new(model.Expense)

	//クエリパラメーターからイベントIDを取得
	event_id, _ := strconv.Atoi(common.GetQueryParam(r))

	//イベントIDを格納
	expense.EventId = event_id

	//名前を格納
	expense.Name = name

	//金額を格納
	expense.Total = total

	//備考を格納
	expense.Remarks = remarks

	fmt.Println("通った")
	//支払い情報を保存
	expense.AddExpense()

	//各参加者の負担金の登録
	//イベントIDから参加者データを取得
	members := model.GetMembers(event_id)

	fmt.Println("通った2")
	//参加者人数を取得
	member_count := len(members)

	//各参加者の負担金を設定
	each_price := total / member_count

	//金額が参加人数で割り切れない場合
	if fraction := total % member_count; fraction != 0 {
		event := new(model.Event)
		event.Id = event_id
		event.UpdatePool()
		fmt.Println("通ったv")
	}
	fmt.Println("通った3")

	fmt.Printf("渡した値：%d,%d,%d", event_id, each_price, expense.Total)

	//イベント参加者各位の負担金額を保存する
	err := model.CreateMemberExpense(event_id, expense, each_price)
	fmt.Println("通った4")
	if err != nil {
		fmt.Println(err)
	}

	RenderTemplate(w, "view/expense/complete", nil)

}

func saveExpenseHandler(w http.ResponseWriter, r *http.Request) {

}

func showExpenseHandler(w http.ResponseWriter, r *http.Request) {

}
