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

func (expense *Expense) AddExpense() {

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

}
