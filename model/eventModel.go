package model

import (
	"crypto/rand"
	"fmt"
	"shareTravel/common"
	"time"
)

type Event struct {
	Id         int
	AuthKey    string
	Pool       int
	Name       string `validate:"required"`
	Date       string `validate:"required"`
	CreateTime time.Time
	UpdateTime time.Time
}

var event_table = "event"

var event_columns = [7]string{"id", "auth_key", "pool", "name", "date", "create_time", "update_time"}

//イベント情報登録関数
func (event *Event) CreateEvent() string {

	//認証キー(ランダム文字列)の取得
	//key, err := createAuthKey()
	key := "0fQasfadsfas"
	var err error
	if err != nil {
		fmt.Println(err)
		return ""
	}

	//認証キーの重複チェック
	for {
		if !DuplicateCheck(event_table, "auth_key", key) {
			key, err = createAuthKey()
			if err != nil {
				fmt.Println(err)
				return ""
			}
			continue
		}
		break
	}

	//認証IDの設定
	event.AuthKey = key

	//現在時刻の取得
	now := time.Now()
	event.CreateTime = now

	//データ登録用のマップを生成
	data := event.EventAutoMapperForModel()

	//DB登録
	err = insert(event_table, data)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	//認証IDを返却
	return key
}

//クエリ作成用マップのオートマッピング関数
func (event *Event) EventAutoMapperForModel() map[string]interface{} {

	//マップデータの初期化
	data := map[string]interface{}{}

	//各プロパティに対して値がセットされている場合、カラム名に紐づけてマッピング
	if event.Id != 0 {
		data[event_columns[0]] = event.Id
	}

	if event.AuthKey != "" {
		data[event_columns[1]] = event.AuthKey
	}

	if event.Pool != 0 {
		data[event_columns[2]] = event.Pool
	}

	if event.Name != "" {
		data[event_columns[3]] = event.Name
	}

	if event.Date != "" {
		data[event_columns[4]] = event.Date
	}

	if !event.CreateTime.IsZero() {
		data[event_columns[5]] = event.CreateTime
	}

	if !event.UpdateTime.IsZero() {
		data[event_columns[6]] = event.UpdateTime
	}

	//マッピングデータを返却
	return data

}

//認証ID生成用関数
func createAuthKey() (string, error) {

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

//
//リファクタリング未済
//
func GetEvent(event *Event) *Event {

	var err error

	if event.Id != 0 {
		err = Db.QueryRow("SELECT id,auth_key,name,date,pool FROM event WHERE id = ?", event.Id).Scan(&event.Id, &event.AuthKey, &event.Name, &event.Date, &event.Pool)
	} else if event.AuthKey != "" {
		err = Db.QueryRow("SELECT id,auth_key,name,date,pool FROM event WHERE auth_key = ?", event.AuthKey).Scan(&event.Id, &event.AuthKey, &event.Name, &event.Date, &event.Pool)
	} else {
		fmt.Println("it has no id and authkey")
		return nil
	}

	if err != nil {
		return nil
	}

	event.Date = common.TimeFormatterHyphen(event.Date)

	return event
}

func (event *Event) UpdatePool(add int) {

	var pool int

	//現在の端数プールを取得
	err := Db.QueryRow("SELECT pool FROM event WHERE id = ?", event.Id).Scan(&pool)

	if err != nil {
		//変数書き換えに失敗した場合終了
		fmt.Println("Exec error in SELECT")
		return
	}

	//レコードの更新処理
	statement := "UPDATE event SET pool = ?,update_time = ? WHERE id = ?;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare error")
		return
	}

	//現在時刻を取得
	t := time.Now()

	//処理終了後データソースを閉じる
	defer stmt.Close()

	//端数をプールに加算
	pool = pool + add

	//SQL実行
	stmt.Exec(pool, t, event.Id)
}

func (event *Event) EditPool(pool int, before_pool int) {

	err := Db.QueryRow("SELECT pool FROM event WHERE id = ?", event.Id).
		Scan(&event.Pool)

	if err != nil {
		fmt.Println(err)
		return
	}

	after_pool := event.Pool - before_pool + pool

	statement := "UPDATE event SET pool = ?,update_time = ? WHERE id = ?"
	stmt, err2 := Db.Prepare(statement)

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	defer stmt.Close()
	stmt.Exec(after_pool, time.Now(), event.Id)
}

func (event *Event) UpdateEvent() {

	statement := "UPDATE event SET auth_key = ? ,name = ?,date = ? WHERE id = ? "

	stmt, err := Db.Prepare(statement)

	if err != nil {
		fmt.Println(err)
		return
	}

	//処理終了後データソースを閉じる
	defer stmt.Close()

	stmt.Exec(event.AuthKey, event.Name, event.Date, event.Id)

}
