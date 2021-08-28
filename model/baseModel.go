package model

import (
	"database/sql"
	"fmt"
	"os"
	"shareTravel/common"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//DBへの接続
func Connect() {
	var err error
	driver := os.Getenv("Driver")
	user := os.Getenv("User")
	pass := os.Getenv("Pass")
	host := os.Getenv("Host")
	port := os.Getenv("Port")
	database := os.Getenv("DataBase")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, database)

	Db, err = sql.Open(driver, connection)
	if err != nil {
		panic(err)
	}

}

//重複チェック
func DuplicateCheck(table string, item string, value interface{}) bool {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE %s = ?", table, item)
	err := Db.QueryRow(query, value).Scan(&id)

	if err == nil {
		return false
	} else {
		fmt.Println(err)
		return true
	}
}

//シンプルなINSERT文の生成処理
func insert(table string, values map[string]interface{}) error {

	query := "INSERT INTO" + common.SPACE

	query += table

	var column_query string

	var value_query string

	index := 0

	for column, value := range values {
		if index < len(values)-1 {
			column_query += column + common.COMMA

			switch value := value.(type) {
			case int:
				val_str := strconv.Itoa(value)
				value_query += val_str + common.COMMA
			case string:
				value_query += "'" + value + "'" + common.COMMA
			case time.Time:
				val_str := value.Format(common.TIME_LAYOUT)
				value_query += "'" + val_str + "'" + common.COMMA
			}
		} else {
			column_query += column
			switch value := value.(type) {
			case int:
				val_str := strconv.Itoa(value)
				value_query += val_str
			case string:
				value_query += "'" + value + "'"
			case time.Time:
				val_str := value.Format(common.TIME_LAYOUT)
				value_query += "'" + val_str + "'"
			}
			break
		}
		index++
	}

	column_query = common.SPACE + "(" + column_query + ")" + common.SPACE
	value_query = "values(" + value_query + ")"

	query += column_query + value_query

	stmt, err := Db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec()

	return nil
}

//DBから一つの結果を取得する際のシンプルなSELECT分の実行
func find(table string, subjects []string, conditions map[string]interface{}) *sql.Row {

	subjects_query := common.EMPTY
	condition_query := common.EMPTY

	for i, subject := range subjects {
		if i < len(subjects)-1 {
			subjects_query += subject + common.COMMA
		} else {
			subjects_query += subject
		}
	}

	index := common.ZERO

	for column, value := range conditions {
		if index < len(conditions)-1 {
			condition_query += column + "="

			switch value := value.(type) {
			case int:
				val_str := strconv.Itoa(value)
				condition_query += val_str + " AND "
			case string:
				condition_query += "'" + value + "' AND "
			case time.Time:
				val_str := value.Format(common.TIME_LAYOUT)
				condition_query += "'" + val_str + "' AND "
			}
		} else {
			condition_query += column + "="
			switch value := value.(type) {
			case int:
				val_str := strconv.Itoa(value)
				condition_query += val_str
			case string:
				condition_query += "'" + value + "'"
			case time.Time:
				val_str := value.Format(common.TIME_LAYOUT)
				condition_query += "'" + val_str + "'"
			}
			break
		}
		index++

	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", subjects_query, table, condition_query)

	row := Db.QueryRow(query)

	return row

}

func MakeQueryU() {

}

func MakeQueryD() {

}
