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
	CreateTime          string
}

func (expense *Expense) AddExpense() int {

	statement := "INSERT INTO expense (event_id,temporarily_member,name,total,remarks,create_time) VALUES(?,?,?,?,?,?)"

	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
	}
	t := time.Now()

	defer stmt.Close()
	stmt.Exec(expense.EventId, expense.TemporarilyMemberId, expense.Name, expense.Total, expense.Remarks, t)

	if err != nil {
		fmt.Println(err)
	}
	var id int
	err2 := Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)

	if err2 != nil {
		return 0
	}

	return id
}

func (expense *Expense) UpdateExpense() {

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

	err := Db.QueryRow("SELECT event_id,temporarily_member,total,remarks,name FROM expense WHERE id = ?", expense.Id).
		Scan(&expense.EventId, &expense.TemporarilyMemberId, &expense.Total, &expense.Remarks, &expense.Name)

	if err != nil {
		fmt.Println(err)
	}

}

func (event *Event) GetExpensesByEventId() []Expense {

	rows, err := Db.Query("SELECT id,total,name,remarks,temporarily_member,create_time from expense WHERE event_id = ?", event.Id)

	//SQLエラー処理
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//DBからの取得データをスライスに変換
	var expenses []Expense

	for rows.Next() {
		expense := Expense{}
		err := rows.Scan(&expense.Id, &expense.Total, &expense.Name, &expense.Remarks, &expense.TemporarilyMemberId, &expense.CreateTime)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		expense.CreateTime = common.TimeFormatter(expense.CreateTime)
		expenses = append(expenses, expense)

	}

	//値の返却
	return expenses
}
