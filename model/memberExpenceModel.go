package model

import (
	"fmt"
	"time"
)

type MemberExpense struct {
	Id         int
	MemberId   int
	ExpenseId  int
	Price      int
	CreateTime time.Time
	UpdateTime time.Time
}

var member_expense_table = "member_expense"
var member_expense_columns = []string{"id", "member_id", "expense_id", "price", "create_time", "update_time"}

func (member_expense *MemberExpense) CreateMemberExpense() error {

	//構造体をマップに変換
	status := member_expense.MemberExpenseAutoMapperForModel()

	//データ保存
	err := insert(member_expense_table, status)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//DB操作用のカラムと値の紐付けを行う
func (member_expense *MemberExpense) MemberExpenseAutoMapperForModel() map[string]interface{} {

	//マップデータの初期化
	data := map[string]interface{}{}

	//各プロパティに対して値がセットされている場合、カラム名に紐づけてマッピング
	if member_expense.Id != 0 {
		data[member_expense_columns[0]] = member_expense.Id
	}

	if member_expense.MemberId != 0 {
		data[member_expense_columns[1]] = member_expense.MemberId
	}

	if member_expense.ExpenseId != 0 {
		data[member_expense_columns[2]] = member_expense.ExpenseId
	}

	if member_expense.Price != 0 {
		data[member_expense_columns[3]] = member_expense.Price
	}

	if !member_expense.CreateTime.IsZero() {
		data[member_expense_columns[4]] = member_expense.CreateTime
	}

	if !member_expense.UpdateTime.IsZero() {
		data[member_expense_columns[5]] = member_expense.UpdateTime
	}

	//マッピングデータを返却
	return data

}

func (member *Member) SearchMemberExpense(expense_id int) {

	err := Db.QueryRow("SELECT price FROM member_expense WHERE member_id = ? AND expense_id = ?", member.Id, expense_id).
		Scan(&member.Calculate)

	if err != nil {
		member.Calculate = 0
	}
}

//
//
//
//
//
//
//
//
//
//
//
//
//

func GetMemberExpensesAll(member_id int) []MemberExpense {

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
			fmt.Println(err)
			panic(err.Error())
		}
		member_expenses = append(member_expenses, member_expense)
	}

	return member_expenses
}

func (member *Member) GetMemberExpense() {

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
			fmt.Println(err)
			panic(err.Error())
		}
		totalExpense += expense
	}

	member.Total = totalExpense
}

func (member_expense *MemberExpense) UpdateMemberExpense() {

	statement := "UPDATE member_expense SET price = ? , update_time = ? WHERE member_id = ? AND expense_id  = ?"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()
	stmt.Exec(member_expense.Price, time.Now(), member_expense.MemberId, member_expense.ExpenseId)
}
