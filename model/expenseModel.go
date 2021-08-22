package model

import (
	"fmt"
	"shareTravel/common"
	"shareTravel/form"
	"time"
)

type Expense form.Expense

func init() {
	OpenSQL()
}

func (expense *Expense) AddExpense() int {

	statement := "INSERT INTO expense (event_id,temporarily_member,name,total,remarks,create_time) VALUES(?,?,?,?,?,?)"

	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare error")
	}
	t := time.Now()

	defer stmt.Close()
	stmt.Exec(expense.EventId, expense.TemporarilyMemberId, expense.Name, expense.Total, expense.Remarks, t)

	if err != nil {
		fmt.Println("Exec error")
	}
	var id int
	err2 := Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)

	if err2 != nil {
		return 0
	}

	return id
}

func (expense *Expense) UpdateExpense() {
	OpenSQL()
	statement := "UPDATE expense SET temporarily_member = ? ,update_time = ? WHERE id = ?"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer stmt.Close()
	stmt.Exec(expense.TemporarilyMemberId, time.Now(), expense.Id)
}

func (expense *Expense) GetExpense() {

	OpenSQL()

	err := Db.QueryRow("SELECT event_id,temporarily_member,total,remarks,name FROM expense WHERE id = ?", expense.Id).
		Scan(&expense.EventId, &expense.TemporarilyMemberId, &expense.Total, &expense.Remarks, &expense.Name)

	if err != nil {
		fmt.Println(err)
	}

}

func (event *Event) GetExpensesByEventId() []Expense {

	//DB接続
	OpenSQL()
	rows, err := Db.Query("SELECT id,total,name,remarks,temporarily_member,create_time from expense WHERE event_id = ?", event.Id)

	//SQLエラー処理
	if err != nil {
		fmt.Println("Query Error")
		return nil
	}

	//DBからの取得データをスライスに変換
	var expenses []Expense

	for rows.Next() {
		expense := Expense{}
		err := rows.Scan(&expense.Id, &expense.Total, &expense.Name, &expense.Remarks, &expense.TemporarilyMemberId, &expense.CreateTime)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		expense.CreateTime = common.TimeFormatter(expense.CreateTime)
		expenses = append(expenses, expense)

	}

	//値の返却
	return expenses
}
