package model

import (
	"fmt"
	"time"
)

func CreateMemberExpense(event_id int, expense *Expense, each_price int) error {

	OpenSQL()

	//イベント参加者を全て取得
	rows, _ := Db.Query("SELECT id FROM member WHERE event_id = ?", event_id)

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
		if member_id != 0 {
			fmt.Println(member_id)
			fmt.Println(members_id)
			members_id = append(members_id, member_id)
		}
	}

	t := time.Now()

	statement := "INSERT INTO member_expense (member_id,event_id,expense_id,price,create_time) VALUES(?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)

	for _, member_id := range members_id {
		_, err := stmt.Exec(member_id, event_id, expense.Id, each_price, t)

		fmt.Println(err)

	}

	defer stmt.Close()

	if err != nil {
		fmt.Println("Exec error")

		return err
	}

	return nil
}
