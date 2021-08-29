package controller

import (
	"net/http"
	"shareTravel/common"
	"shareTravel/model"
	"strconv"
	"strings"
	"time"
)

func ExpenseHandler(w http.ResponseWriter, r *http.Request, path string) {
	arr := strings.Split(path, "/")

	switch arr[0] {
	case ADD:
		addExpenseHandler(w, r)
	case COMPLETE:
		completeExpenseHandler(w, r)
	case CONFIRM:
		confirmExpenseHandler(w, r)
	case CALCULATE:
		calculateExpenseHandler(w, r)
	case EDIT:
		editExpenseHandler(w, r)
	case EDIT_CALCULATE:
		editCalculateHandler(w, r)
	case UPDATE:
		updateHandler(w, r)

	}
}

func addExpenseHandler(w http.ResponseWriter, r *http.Request) {

	event := new(model.Event)
	//クエリパラメータの取得
	event_id := common.GetQueryParam(r)
	event.Id, _ = strconv.Atoi(event_id)

	status := autoMapperForView(event)

	RenderTemplate(w, EXPENSE_ADD_PATH, status)

}

func confirmExpenseHandler(w http.ResponseWriter, r *http.Request) {

	//イベント構造体ポインタ取得
	event := new(model.Event)
	event_id := common.GetQueryParam(r)
	event.Id, _ = strconv.Atoi(event_id)

	//参加者がいない場合エラーハンドリングを行う
	if model.GetMembers(event.Id) == nil {
		errs := map[string]string{}
		errs["Members1"] = `参加者がいない場合、会計の登録はできません。`
		errs["Members2"] = `先に参加者の登録をしてください。`
		status := autoMapperForView(event)
		errorHandler(w, EXPENSE_ADD_PATH, status, errs)
		return
	}

	//入力値から支払い構造体ポインタを作成
	expense, errs := formValueEncodeForExpense(r)

	//入力値エラーがあった場合
	if errs != nil {
		status := autoMapperForView(event, expense)
		errorHandler(w, EXPENSE_ADD_PATH, status, errs)
		return
	}

	members := model.GetMembers(event.Id)

	status := defalutPriceSet(members, expense, event)

	RenderTemplate(w, EXPENSE_CONFIRM_PATH, status)
}

//再計算処理
func calculateExpenseHandler(w http.ResponseWriter, r *http.Request) {

	event := new(model.Event)
	event_id, _ := strconv.Atoi(common.GetQueryParam(r))
	event.Id = event_id
	event.GetEvent()

	//参加者一覧データ処理(文字列型で送信されたデータを構造体スライスに変換)
	members := model.GetMembers(event.Id)

	//hidden値から支払い構造体ポインタを作成
	expense, errs := formValueEncodeForExpense(r)

	//エラーがあった場合
	if errs != nil {

		status := defalutPriceSet(members, expense, event)
		errorHandler(w, EXPENSE_ADD_PATH, status, errs)
		return
	}

	//変更されたメンバーを抽出、変更されなかった
	new_price, changed, err := getChangePriceMembers(event, expense, members, r)

	//エラーが発生した場合
	if err != nil {
		errs := map[string]string{}
		errs["Error"] = "金額は半角数字で入力して下さい"

		status := defalutPriceSet(members, expense, event)

		errorHandler(w, EXPENSE_CONFIRM_PATH, status, errs)
		return
	}

	//参加者負担金の合計
	var member_total int

	//負担金配列を作成
Loop:
	for i := 0; i < len(members); i++ {
		for member_id, change := range changed {
			if member_id == members[i].Id {
				members[i].Calculate = change
				member_total += members[i].Calculate
				continue Loop
			}
		}
		members[i].Calculate = new_price
		member_total += members[i].Calculate
	}

	//支払い合計
	total, _ := strconv.Atoi(r.FormValue("total"))

	//端数
	pool := total - member_total

	expense.Pool = pool

	status := autoMapperForView(event, expense, members)
	status["Price"] = new_price
	status["Temporarily"], _ = strconv.Atoi(r.FormValue("temporarily"))

	if r.FormValue("slash") == "true" {
		status["Slash"] = "ture"
	} else {
		status["Slash"] = "false"
	}

	RenderTemplate(w, "view/expense/confirm", status)

}

func completeExpenseHandler(w http.ResponseWriter, r *http.Request) {

	event := new(model.Event)
	event_id, _ := strconv.Atoi(common.GetQueryParam(r))
	event.Id = event_id

	//支払い情報を設定
	expense, errs := formValueEncodeForExpense(r)
	//エラーがあった場合
	if errs != nil {
		status := autoMapperForView(event, expense)
		errorHandler(w, EXPENSE_ADD_PATH, status, errs)
		return
	}
	expense.EventId = event_id
	expense.TemporarilyMemberId, _ = strconv.Atoi(r.FormValue("temporarily"))
	expense.Pool, _ = strconv.Atoi(r.FormValue("pool"))
	expense.CreateTime = time.Now()

	//TODO トランザクション
	//支払い情報を保存(IDを戻り値として取得)
	var err error
	expense.Id, err = expense.AddExpense()

	//エラー処理
	if err != nil {
		errs := map[string]string{}
		errs["Error"] = "予期せぬエラーが発生しました"

		status := autoMapperForView(event)

		errorHandler(w, EXPENSE_ADD_PATH, status, errs)
		return
	}

	//各参加者の負担金の登録
	//イベントIDから参加者データを取得
	members := model.GetMembers(event_id)

	//各参加者の負担金の保存
	for _, member := range members {
		member_expense := new(model.MemberExpense)
		member_expense.MemberId = member.Id
		member_expense.ExpenseId = expense.Id
		member_expense.Price, _ = strconv.Atoi(r.FormValue(strconv.Itoa(member.Id)))
		member_expense.CreateTime = time.Now()

		err = member_expense.CreateMemberExpense()
		if err != nil {
			errs := map[string]string{}
			errs["Error"] = "予期せぬエラーが発生しました"

			status := autoMapperForView(event)

			errorHandler(w, EXPENSE_ADD_PATH, status, errs)
			return
		}
	}

	//金額が参加人数で割り切れない場合
	pool, _ := strconv.Atoi(r.FormValue("pool"))
	if pool != 0 {
		event := new(model.Event)
		event.Id = event_id
		err = event.UpdatePool(pool)

		if err != nil {
			errs := map[string]string{}
			errs["Error"] = "予期せぬエラーが発生しました"

			status := autoMapperForView(event)

			errorHandler(w, EXPENSE_ADD_PATH, status, errs)
			return
		}
	}

	//テンプレートの読み込み
	RenderTemplate(w, EXPENSE_COMPLETE_PATH, expense)

}

//リクエストをもとに入力値をイベント用構造体に変換
func formValueEncodeForExpense(r *http.Request) (*model.Expense, map[string]string) {
	expense := new(model.Expense)
	errs := map[string]string{}

	//POST値が存在するかの判別
	if name := r.FormValue("name"); name != EMPTY {
		expense.Name = name
	} else {
		errs["Name"] = "名前は必須入力です。"
	}

	if price := r.FormValue("total"); price != EMPTY {
		var err error
		expense.Total, err = strconv.Atoi(price)
		if err != nil {
			errs["Price"] = "合計金額は半角数字で入力してください。"
		}
	} else {
		errs["Price"] = "合計金額は必須入力です。"
	}

	if remarks := r.FormValue("remarks"); remarks != EMPTY {
		expense.Remarks = remarks
	}

	if len(errs) != 0 {
		return expense, errs
	} else {
		return expense, nil
	}
}

func defalutPriceSet(members []*model.Member, expense *model.Expense, event *model.Event) map[string]interface{} {
	//合計金額を人数で割った値を算出
	member_count := len(members)
	default_price := (expense.Total / member_count) / 100 * 100

	//初期負担金額をセット
	for i := 0; i < len(members); i++ {
		members[i].Calculate = default_price
	}

	//端数を計算
	default_pool := expense.Total - (default_price * len(members))
	expense.Pool = default_pool

	//値のセット
	status := autoMapperForView(event, expense, members)
	status["Price"] = &default_price
	status["Temporarily"] = members[0].Id

	return status
}

//再計算時に金額を変更された参加者を抽出、変更後の差し引き合計金額を計算
func getChangePriceMembers(event *model.Event, expense *model.Expense, members []*model.Member, r *http.Request) (int, map[int]int, error) {

	changed := map[int]int{}
	total := expense.Total

	for _, member := range members {
		//入力値とデフォルトの負担金額が異なる場合（金額の変更があった場合）
		if r.FormValue(strconv.Itoa(member.Id)) != r.FormValue("price") {
			//入力値を配列に格納し、合計金額を減算
			price, err := strconv.Atoi(r.FormValue(strconv.Itoa(member.Id)))
			if err != nil {
				return 0, nil, err
			}
			changed[member.Id] = price
			total = total - price
		}
	}

	new_price := 0

	if len(members)-len(changed) != 0 {
		//再計算（変更されなかった参加者の負担金）
		new_price = total / (len(members) - len(changed))
	}

	//切り捨てチェックが入っていた場合計算値を切り捨て
	if r.FormValue("slash") == "true" {
		new_price = (new_price / 100) * 100
	}

	//計算値の返却
	return new_price, changed, nil
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

	var member_price_total int

	for _, member := range members {
		member_price_total += member.Calculate
	}

	status := autoMapperForView(event, expense, members)
	status["Slash"] = "true"

	autoMapperForView(event, expense, members)

	//テンプレートの読み込み
	RenderTemplate(w, EXPENSE_EDIT_PATH, status)
}

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

func editCalculateHandler(w http.ResponseWriter, r *http.Request) {

	event_id, _ := strconv.Atoi(r.FormValue("event"))
	event := new(model.Event)
	event.Id = event_id
	event.GetEvent()

	//参加者一覧データ処理(文字列型で送信されたデータを構造体スライスに変換)
	members := model.GetMembers(event.Id)
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
