package controller

import (
	"fmt"
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

	default_price := defalutPriceSet(members, expense, event)

	//値のセット
	status := autoMapperForView(event, expense, members)
	status["Price"] = &default_price
	status["Temporarily"] = members[0].Id

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

		default_price := defalutPriceSet(members, expense, event)
		status := autoMapperForView(event, expense, members)
		status["Price"] = &default_price
		status["Temporarily"] = members[0].Id
		errorHandler(w, EXPENSE_ADD_PATH, status, errs)
		return
	}

	changed := map[int]int{}
	total := expense.Total

	for _, member := range members {

		//入力値とデフォルトの負担金額が異なる場合（金額の変更があった場合）
		if r.FormValue(strconv.Itoa(member.Id)) != r.FormValue("price") {
			//入力値を配列に格納し、合計金額を減算
			price, err := strconv.Atoi(r.FormValue(strconv.Itoa(member.Id)))
			if err != nil {
				errs := map[string]string{}
				errs["Error"] = "金額は半角数字で入力して下さい"

				default_price := defalutPriceSet(members, expense, event)
				status := autoMapperForView(event, expense, members)
				status["Price"] = &default_price
				status["Temporarily"] = members[0].Id

				errorHandler(w, EXPENSE_CONFIRM_PATH, status, errs)
				return
			}
			changed[member.Id] = price
			total = total - price
		}
	}

	//変更があった参加者を除いた参加者の負担額を再計算
	new_price := getNewPrice(members, r, changed, total)

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

	//端数
	pool := expense.Total - member_total

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

//支払い登録完了ハンドラー
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

//支払い情報の編集ハンドラ
func editExpenseHandler(w http.ResponseWriter, r *http.Request) {

	//支払い情報設定
	expense := new(model.Expense)
	expense.Id, _ = strconv.Atoi(common.GetQueryParam(r))
	expense.GetExpense()

	//イベント情報の設定
	event := new(model.Event)
	event.Id = expense.EventId
	event.GetEvent()

	//参加者の取得
	members := model.GetMembers(event.Id)

	//参加者に紐づく会計情報の取得
	for i := 0; i < len(members); i++ {
		members[i].SearchMemberExpense(expense.Id)
	}

	status := autoMapperForView(event, expense, members)
	status["Slash"] = "true"
	status["BeforePool"] = expense.Pool

	//テンプレートの読み込み
	RenderTemplate(w, EXPENSE_EDIT_PATH, status)
}

func editCalculateHandler(w http.ResponseWriter, r *http.Request) {

	//イベント取得
	event := new(model.Event)
	event.Id, _ = strconv.Atoi(r.FormValue("event"))
	event.GetEvent()

	//参加者一覧データ処理(文字列型で送信されたデータを構造体スライスに変換)
	members := model.GetMembers(event.Id)

	//支払い情報を設定
	expense, errs := formValueEncodeForExpense(r)
	//エラーがあった場合
	if errs != nil {
		status := autoMapperForView(event, expense)
		errorHandler(w, EXPENSE_EDIT_PATH, status, errs)
		return
	}
	expense.Id, _ = strconv.Atoi(common.GetQueryParam(r))
	expense.EventId = event.Id
	expense.TemporarilyMemberId, _ = strconv.Atoi(r.FormValue("temporarily"))
	expense.Pool, _ = strconv.Atoi(r.FormValue("pool"))

	//負担金額に変更があった参加者マップ及び、合計金額との差分を作成
	changed := map[int]int{}
	total := expense.Total

	//変更されたメンバーを抽出
	for _, member := range members {
		if r.FormValue(strconv.Itoa(member.Id)) != r.FormValue("before"+strconv.Itoa(member.Id)) {
			price, err := strconv.Atoi(r.FormValue(strconv.Itoa(member.Id)))

			//エラー処理
			if err != nil {
				fmt.Println(err)

				//参加者に紐づく会計情報の取得
				for i := 0; i < len(members); i++ {
					members[i].SearchMemberExpense(expense.Id)
				}

				status := autoMapperForView(event, expense, members)
				status["BeforePool"], _ = strconv.Atoi(r.FormValue("beforePool"))

				if r.FormValue("slash") == "true" {
					status["Slash"] = "ture"
				} else {
					status["Slash"] = "false"
				}

				errs := map[string]string{}
				errs["Error"] = "金額は半角数字で入力してください。"

				errorHandler(w, EXPENSE_EDIT_PATH, status, errs)
				return
			}
			changed[member.Id] = price
			total = total - price
		}
	}

	//負担金の変更があった参加者以外の負担金の再計算
	new_price := getNewPrice(members, r, changed, total)

	//参加者の負担金の合計金額を設定及び、参加者構造体ポインタに負担金情報を設定
	member_total := 0
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

	//端数を支払い情報に設定
	pool := expense.Total - member_total
	expense.Pool = pool

	//画面表示情報成形処理
	status := autoMapperForView(event, expense, members)
	status["BeforePool"], _ = strconv.Atoi(r.FormValue("beforePool"))

	if r.FormValue("slash") == "true" {
		status["Slash"] = "ture"
	} else {
		status["Slash"] = "false"
	}

	//テンプレートの読み込み
	RenderTemplate(w, EXPENSE_EDIT_PATH, status)

}

func updateHandler(w http.ResponseWriter, r *http.Request) {

	//イベント取得
	event := new(model.Event)
	event.Id, _ = strconv.Atoi(r.FormValue("event"))
	event.GetEvent()

	//支払い情報を設定
	expense, errs := formValueEncodeForExpense(r)
	//エラーがあった場合
	if errs != nil {
		status := autoMapperForView(event, expense)
		errorHandler(w, EXPENSE_ADD_PATH, status, errs)
		return
	}
	expense.Id, _ = strconv.Atoi(common.GetQueryParam(r))
	expense.EventId = event.Id
	expense.TemporarilyMemberId, _ = strconv.Atoi(r.FormValue("temporarily"))
	expense.Pool, _ = strconv.Atoi(r.FormValue("pool"))

	//支払い情報を保存(IDを戻り値として取得)
	expense.UpdateExpense()

	//各参加者の負担金の登録
	//イベントIDから参加者データを取得
	members := model.GetMembers(event.Id)

	//各参加者の負担金の保存
	for _, member := range members {
		member_expense := new(model.MemberExpense)
		member_expense.MemberId = member.Id
		member_expense.ExpenseId = expense.Id
		member_expense.Price, _ = strconv.Atoi(r.FormValue(strconv.Itoa(member.Id)))
		member_expense.UpdateMemberExpense()
	}

	//金額が参加人数で割り切れない場合
	pool, _ := strconv.Atoi(r.FormValue("pool"))
	before_pool, _ := strconv.Atoi(r.FormValue("beforePool"))

	//TODO
	//端数の更新処理
	if pool != before_pool {
		event := new(model.Event)
		event.Id, _ = strconv.Atoi(r.FormValue("event"))
		event.EditPool(pool, before_pool)
	}

	//テンプレートの読み込み
	RenderTemplate(w, EXPENSE_UPDATE_PTH, expense)

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

//初期表示用の各参加者負担金額の計算
func defalutPriceSet(members []*model.Member, expense *model.Expense, event *model.Event) int {
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

	return default_price

}

//再計算時に金額を変更された参加者を抽出、変更後の差し引き合計金額を計算
func getNewPrice(members []*model.Member, r *http.Request, changed map[int]int,
	total int) int {

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
	return new_price
}
