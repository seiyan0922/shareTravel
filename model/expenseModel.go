package model

import (
	"fmt"
	"shareTravel/common"
	"time"
)

type Expense struct {
	Id                  int
	EventId             int
	TemporarilyMemberId int
	Name                string
	Total               int
	Remarks             string
	Pool                int
	DateStr             string
	CreateTime          time.Time
	UpdateTime          time.Time
}

var expense_table = "expense"

var expense_columns = []string{"id", "event_id", "temporarily_member", "total", "remarks", "name", "pool", "create_time", "update_time"}

//参加者の建て替え金額を取得する
func (member *Member) GetMemberTemporarily() error {

	select_colomns := []string{expense_columns[3]}

	expense := new(Expense)
	expense.TemporarilyMemberId = member.Id

	status := expense.ExpenseAutoMapperForModel()

	rows, err := findAll(expense_table, select_colomns, status)

	if err != nil {
		fmt.Println(err)
		return err
	}

	var total_temporarily int

	for rows.Next() {
		var temporarily int
		err := rows.Scan(&temporarily)
		if err != nil {
			fmt.Println(err)
			return err
		}
		total_temporarily += temporarily
	}

	member.Temporarily = total_temporarily

	return nil
}

//イベントに紐づく支払い一覧を取得
func (event *Event) GetExpenses() []*Expense {

	expense := new(Expense)
	expense.EventId = event.Id

	condition := expense.ExpenseAutoMapperForModel()
	select_columns := []string{expense_columns[0], expense_columns[2], expense_columns[3], expense_columns[4], expense_columns[5], expense_columns[7]}
	rows, err := findAll(expense_table, select_columns, condition)

	//SQLエラー処理
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//DBからの取得データをスライスに変換
	var expenses []*Expense

	for rows.Next() {
		expense := new(Expense)
		err := rows.Scan(&expense.Id, &expense.TemporarilyMemberId, &expense.Total, &expense.Remarks, &expense.Name, &expense.CreateTime)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		expense.DateStr = common.TimeFormatter(expense.CreateTime)
		expenses = append(expenses, expense)
	}

	//値の返却
	return expenses
}

//DB操作用のカラムと値の紐付けを行う
func (expense *Expense) ExpenseAutoMapperForModel() map[string]interface{} {

	//マップデータの初期化
	data := map[string]interface{}{}

	//各プロパティに対して値がセットされている場合、カラム名に紐づけてマッピング
	if expense.Id != 0 {
		data[expense_columns[0]] = expense.Id
	}

	if expense.EventId != 0 {
		data[expense_columns[1]] = expense.EventId
	}

	if expense.TemporarilyMemberId != 0 {
		data[expense_columns[2]] = expense.TemporarilyMemberId
	}

	if expense.Total != 0 {
		data[expense_columns[3]] = expense.Total
	}

	if expense.Remarks != "" {
		data[expense_columns[4]] = expense.Remarks
	}

	if expense.Name != "" {
		data[expense_columns[5]] = expense.Name
	}

	if expense.Pool != 0 {
		data[expense_columns[6]] = expense.Pool
	}

	if !expense.CreateTime.IsZero() {
		data[expense_columns[7]] = expense.CreateTime
	}

	if !expense.UpdateTime.IsZero() {
		data[expense_columns[8]] = expense.UpdateTime
	}

	//マッピングデータを返却
	return data

}

func (expense *Expense) AddExpense() (int, error) {

	status := expense.ExpenseAutoMapperForModel()

	err := insert(expense_table, status)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var id int
	err = Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

//
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
//
//

func (expense *Expense) UpdateExpense() {

	statement := "UPDATE expense SET temporarily_member = ? ,update_time = ? WHERE id = ?"
	stmt, err := Db.Prepare(statement)

	fmt.Println(expense.TemporarilyMemberId)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer stmt.Close()
	stmt.Exec(expense.TemporarilyMemberId, time.Now(), expense.Id)

	fmt.Println(expense.Id)
}

func (expense *Expense) GetExpense() {

	select_columns := []string{expense_columns[1], expense_columns[2], expense_columns[3], expense_columns[4],
		expense_columns[5], expense_columns[6]}
	status := expense.ExpenseAutoMapperForModel()
	row := find(expense_table, select_columns, status)

	err := row.Scan(&expense.EventId, &expense.TemporarilyMemberId, &expense.Total, &expense.Remarks, &expense.Name, &expense.Pool)

	if err != nil {
		fmt.Println(err)
	}

}
