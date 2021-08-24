package model

import (
	"fmt"
	"shareTravel/form"
	"time"
)

type MemberExpense form.MemberExpense

func (member_expense *MemberExpense) CreateMemberExpense() {

	//DB接続
	OpenSQL()
	defer Db.Close()

	t := time.Now()

	//クエリの作成
	statement := "INSERT INTO member_expense (member_id,expense_id,price,create_time) VALUES(?,?,?,?)"

	//実行準備
	stmt, err := Db.Prepare(statement)

	//クエリ実行
	_, err2 := stmt.Exec(member_expense.MemberId, member_expense.ExpenseId, member_expense.Price, t)

	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	if err2 != nil {
		fmt.Println("Exec error")
	}
}

func GetMemberExpensesAll(member_id int) []MemberExpense {
	//DB接続
	OpenSQL()
	defer Db.Close()

	//クエリ実行
	rows, err := Db.Query("SELECT expense_id, price from member_expense WHERE member_id = ?", member_id)

	if err != nil {
		fmt.Println(err)
	}

	member_expenses := []MemberExpense{}

	for rows.Next() {
		var member_expense MemberExpense
		err := rows.Scan(&member_expense.ExpenseId, &member_expense.Price)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		member_expenses = append(member_expenses, member_expense)
	}

	return member_expenses
}

func GetMemberExpense(member *Member) {

	//DB接続
	OpenSQL()
	defer Db.Close()

	//クエリ実行
	rows, err := Db.Query("SELECT price from member_expense WHERE member_id = ?", member.Id)

	if err != nil {
		fmt.Println(err)
	}

	totalExpense := 0

	for rows.Next() {
		var expense int
		err := rows.Scan(&expense)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		totalExpense += expense
	}

	member.Total = totalExpense
}

func (member *Member) SearchMemberExpense(expense_id int) {
	OpenSQL()
	defer Db.Close()

	err := Db.QueryRow("SELECT price FROM member_expense WHERE member_id = ? AND expense_id = ?", member.Id, expense_id).
		Scan(&member.Calculate)

	if err != nil {
		member.Calculate = 0
	}
}

func (member_expense *MemberExpense) UpdateMemberExpense() {
	OpenSQL()
	defer Db.Close()
	statement := "UPDATE member_expense SET price = ? , update_time = ? WHERE member_id = ? AND expense_id  = ?"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()
	stmt.Exec(member_expense.Price, time.Now(), member_expense.MemberId, member_expense.ExpenseId)
}
