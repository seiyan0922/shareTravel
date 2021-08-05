package model

import (
	"fmt"
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
