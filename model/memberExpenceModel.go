package model

import (
	"fmt"
	"shareTravel/form"
	"time"
)

type MemberExpense form.MemberExpense

func (member_expense *MemberExpense) CreateMemberExpense() {

	OpenSQL()

	t := time.Now()

	statement := "INSERT INTO member_expense (member_id,expense_id,price,create_time) VALUES(?,?,?,?)"
	stmt, err := Db.Prepare(statement)

	_, err2 := stmt.Exec(member_expense.MemberId, member_expense.ExpenseId, member_expense.Price, t)

	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	if err2 != nil {
		fmt.Println("Exec error")
	}
}
