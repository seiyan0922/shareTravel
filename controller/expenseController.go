package controller

import (
	"net/http"
	"shareTravel/common"
	"shareTravel/model"
	"strconv"
	"strings"
)

//リファクタリング未済

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
	case "edit":
		editExpenseHandler(w, r)
	case "editcalculate":
		editCalculateHandler(w, r)
	case "update":
		updateHandler(w, r)

	}
}

func addExpenseHandler(w http.ResponseWriter, r *http.Request) {

	//クエリパラメータの取得
	event_id := common.GetQueryParam(r)
	event := new(model.Event)
	event.Id, _ = strconv.Atoi(event_id)

	RenderTemplate(w, "view/expense/add", event)

}

func confirmExpenseHandler(w http.ResponseWriter, r *http.Request) {

	event_id := common.GetQueryParam(r)
	event := new(model.Event)
	event.Id, _ = strconv.Atoi(event_id)

	if model.GetMembers(event.Id) == nil {

		errs := map[string]string{}
		errs["Members1"] = `参加者がいない場合会計の登録はできません。`
		errs["Members2"] = `先に参加者の登録をしてください。`
		status := autoMapperForView(event)
		errorHandler(w, "view/expense/add", status, errs)
		return
	}

	expense := new(model.Expense)
	expense.EventId = event.Id
	expense.Name = r.FormValue("name")
	expense.Total, _ = strconv.Atoi(r.FormValue("price"))
	expense.Remarks = r.FormValue("remarks")

	members := model.GetMembers(event.Id)
	member_count := len(members)

	default_price := (expense.Total / member_count) / 100 * 100
	default_pool := expense.Total - (default_price * len(members))

	//値のセット
	status := make(map[string]interface{})
	status["Event"] = &event
	status["Expense"] = &expense
	status["Members"] = &members
	status["MemberCount"] = &member_count
	status["Price"] = &default_price
	status["Pool"] = &default_pool

	RenderTemplate(w, "view/expense/confirm", status)
}

func calculateExpenseHandler(w http.ResponseWriter, r *http.Request) {

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

	if r.FormValue("slash") == "true" {
		new_price = (new_price / 100) * 100
	}

	new_price_map := map[int]int{}

	//負担金配列を作成
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

	status := make(map[string]interface{})

	event_id, _ := strconv.Atoi(r.FormValue("event"))
	event := new(model.Event)
	event.Id = event_id
	event.GetEvent()

	var total_price int

	for _, price := range new_price_map {
		total_price += price
	}

	total2, _ := strconv.Atoi(r.FormValue("total"))

	pool := total2 - total_price

	status["Event"] = event
	status["Expense"] = expense
	status["Members"] = members
	status["Pool"] = pool
	status["Price"] = new_price
	status["Temporarily"], _ = strconv.Atoi(r.FormValue("temporarily"))

	if r.FormValue("slash") == "true" {
		status["Slash"] = "ture"
	} else {
		status["Slash"] = "false"
	}

	RenderTemplate(w, "view/expense/calculate", status)

}

func completeExpenseHandler(w http.ResponseWriter, r *http.Request) {

	//POSTデータ
	total, _ := strconv.Atoi(r.FormValue("total"))
	name := r.FormValue("name")
	remarks := r.FormValue("remarks")

	expense := new(model.Expense)

	event_id, _ := strconv.Atoi(common.GetQueryParam(r))
	expense.EventId = event_id
	expense.Name = name
	expense.TemporarilyMemberId, _ = strconv.Atoi(r.FormValue("temporarily"))
	expense.Total = total
	expense.Remarks = remarks

	//支払い情報を保存(IDを戻り値として取得)
	expense.Id = expense.AddExpense()

	//各参加者の負担金の登録
	//イベントIDから参加者データを取得
	members := model.GetMembers(event_id)

	//各参加者の負担金の保存
	for _, member := range members {
		member_expense := new(model.MemberExpense)
		member_expense.MemberId = member.Id
		member_expense.ExpenseId = expense.Id
		member_expense.Price, _ = strconv.Atoi(r.FormValue(strconv.Itoa(member.Id)))
		member_expense.CreateMemberExpense()
	}

	//金額が参加人数で割り切れない場合
	pool, _ := strconv.Atoi(r.FormValue("Pool"))
	if pool != 0 {
		event := new(model.Event)
		event.Id = event_id
		event.UpdatePool(pool)
	}

	//テンプレートの読み込み
	RenderTemplate(w, "view/expense/complete", expense)

}

func editExpenseHandler(w http.ResponseWriter, r *http.Request) {

	expense := new(model.Expense)
	expense.Id, _ = strconv.Atoi(common.GetQueryParam(r))

	expense.GetExpense()

	event := new(model.Event)
	event.Id = expense.EventId
	event.GetEvent()

	members := model.GetMembers(event.Id)

	for i := 0; i < len(members); i++ {
		members[i].SearchMemberExpense(expense.Id)
	}

	var pool int
	var member_price_total int

	for _, member := range members {
		member_price_total += member.Calculate
	}

	pool = expense.Total - member_price_total

	status := make(map[string]interface{})
	status["Event"] = &event
	status["Expense"] = &expense
	status["Members"] = &members
	status["Pool"] = &pool
	status["BeforePool"] = &pool
	status["Slash"] = "true"

	//テンプレートの読み込み
	RenderTemplate(w, "view/expense/edit", status)
}

func editCalculateHandler(w http.ResponseWriter, r *http.Request) {

	//参加者一覧データ処理(文字列型で送信されたデータを構造体スライスに変換)
	members := postMembersCnv(r.FormValue("members"))
	expense := postExpenseCnv(r.FormValue("expense"))

	count := 0
	changed := map[int]int{}
	total, _ := strconv.Atoi(r.FormValue("total"))

	//変更されたメンバーを抽出
	for i := 0; i < len(members); i++ {
		if r.FormValue(strconv.Itoa(members[i].Id)) != r.FormValue("before"+strconv.Itoa(members[i].Id)) {
			count++
			price, _ := strconv.Atoi(r.FormValue(strconv.Itoa(members[i].Id)))
			members[i].Calculate = price
			changed[members[i].Id] = price
			total = total - price
		}
	}

	//再計算
	new_price := total / (len(members) - count)

	if r.FormValue("slash") == "true" {
		new_price = (new_price / 100) * 100
	}

	//負担金配列を作成
	for i := 0; i < len(members); i++ {
		if r.FormValue(strconv.Itoa(members[i].Id)) == r.FormValue("before"+strconv.Itoa(members[i].Id)) {
			members[i].Calculate = new_price
		}
	}

	status := make(map[string]interface{})

	event_id, _ := strconv.Atoi(r.FormValue("event"))
	event := new(model.Event)
	event.Id = event_id
	event.GetEvent()

	var total_price int

	for _, member := range members {
		total_price += member.Calculate
	}

	total2, _ := strconv.Atoi(r.FormValue("total"))

	pool := total2 - total_price

	status["Event"] = event
	status["Expense"] = expense
	status["Members"] = members
	status["Pool"] = pool
	status["BeforePool"], _ = strconv.Atoi(r.FormValue("before_pool"))
	status["Temporarily"], _ = strconv.Atoi(r.FormValue("temporarily"))

	if r.FormValue("slash") == "true" {
		status["Slash"] = "ture"
	} else {
		status["Slash"] = "false"
	}

	//テンプレートの読み込み
	RenderTemplate(w, "view/expense/edit", status)

}

func updateHandler(w http.ResponseWriter, r *http.Request) {

	expense := postExpenseCnv(r.FormValue("expense"))

	//支払い情報を保存(IDを戻り値として取得)
	expense.UpdateExpense()

	//各参加者の負担金の登録
	//イベントIDから参加者データを取得
	members := model.GetMembers(expense.EventId)

	//各参加者の負担金の保存
	for _, member := range members {
		member_expense := new(model.MemberExpense)
		member_expense.MemberId = member.Id
		member_expense.ExpenseId = expense.Id
		member_expense.Price, _ = strconv.Atoi(r.FormValue(strconv.Itoa(member.Id)))
		member_expense.UpdateMemberExpense()
	}

	//金額が参加人数で割り切れない場合
	pool, _ := strconv.Atoi(r.FormValue("Pool"))
	before_pool, _ := strconv.Atoi(r.FormValue("BeforePool"))

	if pool != 0 && before_pool != 0 {
		event := new(model.Event)
		event.Id, _ = strconv.Atoi(r.FormValue("event"))
		event.EditPool(pool, before_pool)
	}

	//テンプレートの読み込み
	RenderTemplate(w, "view/expense/update", expense)

}

//POSTデータの変換処理
func postExpenseCnv(str_expense string) model.Expense {

	replaced1 := strings.Replace(str_expense, "[", "", -1)
	replaced2 := strings.Replace(replaced1, "]", "", -1)
	replaced3 := strings.Replace(replaced2, "{", "", -1)

	expenses_arr := strings.Split(replaced3, "} ")

	var expense model.Expense

	for _, str_expense := range expenses_arr {
		expense_arr := strings.Split(str_expense, " ")

		expense_id, _ := strconv.Atoi(expense_arr[0])
		event_id, _ := strconv.Atoi(expense_arr[1])
		temporarilyMemberId, _ := strconv.Atoi(expense_arr[2])
		name := expense_arr[3]
		total, _ := strconv.Atoi(expense_arr[4])
		remarks := expense_arr[5]

		expense.Id = expense_id
		expense.EventId = event_id
		expense.TemporarilyMemberId = temporarilyMemberId
		expense.Name = name
		expense.Total = total
		expense.Remarks = remarks
	}

	return expense
}
