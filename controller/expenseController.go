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
	case "confirm":
		confirmExpenseHandler(w, r)
	case "calculate":
		calculateExpenseHandler(w, r)

	}
}

func addExpenseHandler(w http.ResponseWriter, r *http.Request) {

	event_id := common.GetQueryParam(r)
	event := new(form.Event)
	event.Id, _ = strconv.Atoi(event_id)
	RenderTemplate(w, "view/expense/add", event)

}

func confirmExpenseHandler(w http.ResponseWriter, r *http.Request) {
	event_id := common.GetQueryParam(r)
	event := new(form.Event)
	event.Id, _ = strconv.Atoi(event_id)

	expense := new(model.Expense)
	expense.EventId = event.Id
	expense.Name = r.FormValue("name")
	expense.Total, _ = strconv.Atoi(r.FormValue("price"))
	expense.Remarks = r.FormValue("remarks")

	members := model.GetMembers(event.Id)
	member_count := len(members)

	default_price := expense.Total / member_count
	default_pool := expense.Total % member_count

	status := make(map[string]interface{})
	status["Event"] = event
	status["Expense"] = expense
	status["Members"] = members
	status["MemberCount"] = member_count
	status["Price"] = default_price
	status["Pool"] = default_pool

	RenderTemplate(w, "view/expense/confirm", status)
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
		event.UpdatePool(fraction)
	}

	//イベント参加者各位の負担金額を保存する
	err := model.CreateMemberExpense(event_id, expense, each_price)
	if err != nil {
		fmt.Println(err)
	}

	//テンプレートの読み込み
	RenderTemplate(w, "view/expense/complete", expense)

}

func calculateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	// result := []int{}

	//参加者一覧データ処理(文字列型で送信されたデータを構造体スライスに変換)
	members := postMembersCnv(r.FormValue("members"))
	expense := postExpenseCnv(r.FormValue("expense"))

	count := 0
	changed := map[int]int{}
	total, _ := strconv.Atoi(r.FormValue("total"))

	//変更されたメンバーを抽出
	for _, member := range members {
		if r.FormValue(strconv.Itoa(member.Id)) != r.FormValue("price") {
			count++
			price, _ := strconv.Atoi(r.FormValue(strconv.Itoa(member.Id)))
			changed[member.Id] = price
			total = total - price
		}
	}

	//再計算
	new_price := total / (len(members) - count)

	new_price_map := map[int]int{}

	//新しい負担金配列を作成
Loop:
	for _, member := range members {
		for member_id, change := range changed {
			if member_id == member.Id {
				new_price_map[member_id] = change
				continue Loop
			}
		}
		new_price_map[member.Id] = new_price
	}

	for i, p := range new_price_map {
		for i2 := 0; i2 < len(members); i2++ {
			if members[i2].Id == i {
				members[i2].Calculate = p
			}
		}
	}

	fmt.Println(new_price_map)

	status := make(map[string]interface{})

	event_id, _ := strconv.Atoi(r.FormValue("event"))
	event := new(model.Event)
	event.Id = event_id
	event = model.GetEvent(event)

	var total_price int

	for _, price := range new_price_map {
		total_price += price
	}

	total2, _ := strconv.Atoi(r.FormValue("total"))
	fmt.Println(total_price)
	fmt.Println(total2)

	pool := total2 - total_price
	fmt.Println(pool)

	status["Event"] = event
	status["Expense"] = expense
	status["Members"] = members
	status["Pool"] = pool
	status["Price"] = new_price

	RenderTemplate(w, "view/expense/calculate", status)
}

//POSTデータの変換処理
func postExpenseCnv(str_expense string) form.Expense {
	replaced1 := strings.Replace(str_expense, "[", "", -1)
	replaced2 := strings.Replace(replaced1, "]", "", -1)
	replaced3 := strings.Replace(replaced2, "{", "", -1)

	expenses_arr := strings.Split(replaced3, "} ")

	var expense form.Expense

	for _, str_expense := range expenses_arr {
		expense_arr := strings.Split(str_expense, " ")

		expense_id, _ := strconv.Atoi(expense_arr[0])
		event_id, _ := strconv.Atoi(expense_arr[1])
		name := expense_arr[2]
		total, _ := strconv.Atoi(expense_arr[3])
		remarks := expense_arr[4]

		expense.Id = expense_id
		expense.EventId = event_id
		expense.Name = name
		expense.Total = total
		expense.Remarks = remarks
	}

	return expense
}
