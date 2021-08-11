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

	statement := "INSERT INTO expense (event_id,name,total,remarks,create_time) VALUES(?,?,?,?,?)"

	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare error")
	}
	t := time.Now()

	defer stmt.Close()
	stmt.Exec(expense.EventId, expense.Name, expense.Total, expense.Remarks, t)

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

func (event *Event) GetExpensesByEventId() []Expense {

	//DB接続
	OpenSQL()
	rows, err := Db.Query("SELECT total,name,remarks,create_time from expense WHERE event_id = ?", event.Id)

	//SQLエラー処理
	if err != nil {
		fmt.Println("Query Error")
		return nil
	}

	//DBからの取得データをスライスに変換
	var expenses []Expense

	for rows.Next() {
		expense := Expense{}
		err := rows.Scan(&expense.Total, &expense.Name, &expense.Remarks, &expense.CreateTime)
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
