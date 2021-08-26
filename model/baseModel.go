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

	map_len := len(values)

	index := 0

	for column, value := range values {
		if index < map_len-1 {
			column_query += column + ","

			switch value := value.(type) {
			case int:
				val_str := strconv.Itoa(value)
				value_query += val_str + ","
			case string:
				value_query += "'" + value + "',"
			case time.Time:
				val_str := value.Format(common.TIME_LAYOUT)
				value_query += "'" + val_str + "',"
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

func MakeQueryR() {

}

func MakeQueryU() {

}

func MakeQueryD() {

}
