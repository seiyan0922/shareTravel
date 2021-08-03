package model

import (
	"fmt"
	"time"
)

func CreateMemberExpense(event_id int, expense *Expense, each_price int) error {

	OpenSQL()

	//イベント参加者を全て取得
	rows, _ := Db.Query("SELECT id FROM member WHERE event_id = ?", event_id)

	fmt.Println("通った10")
	fmt.Println(expense.Id)

	var members_id []int

	for rows.Next() {

		//データの一時保存先
		var member_id int

		//データの書き込み
		err := rows.Scan(&member_id)
		if err != nil {
			fmt.Println("Scan error")
			return err
		}
		members_id = append(members_id, member_id)
	}
	fmt.Println("通った11")

	t := time.Now()

	statement := "INSERT INTO member_expense (member_id,event_id,expense_id,price,craete_time) VALUES(?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()

	fmt.Println("通った13")

	for member_id := range members_id {
		fmt.Printf("何かがおかしいメンバーID：%d、イベントID：%d、金額ID:%d,個別の負担：%d", member_id, event_id, expense.Id, each_price)

		_, err := stmt.Exec(member_id, event_id, expense.Id, each_price, t)

		fmt.Println(err)

	}

	if err != nil {
		fmt.Println("Exec error")

		return err
	}
	fmt.Println("通った14")

	return nil
}
