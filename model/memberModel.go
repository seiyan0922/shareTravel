package model

import (
	"fmt"
	"time"
)

type Member struct {
	Id          int
	EventId     int
	Name        string `validate:"required,max=12"`
	Temporarily int
	Calculate   int
	Total       int
	CreateTime  time.Time
	UpdateTime  time.Time
}

var member_table = "member"

var member_columns = []string{"id", "event_id", "name", "create_time", "update_time"}

func (member *Member) SaveMember() error {

	//作成時間を設定
	create_time := time.Now()
	member.CreateTime = create_time

	data := member.MemberAutoMapperForModel()

	err := insert(member_table, data)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetMembers(event_id int) []*Member {

	//イベントポインタ
	event := new(Event)

	//IDを設定
	event.Id = event_id

	member := new(Member)

	member.EventId = (event_id)

	//DB用のマッピングを作成
	status := member.MemberAutoMapperForModel()

	select_columns := []string{member_columns[0], member_columns[2]}

	rows, err := findAll(member_table, select_columns, status)

	if err != nil {
		fmt.Println(err)
	}

	var members []*Member

	//参加者一覧の設定
	for rows.Next() {
		member := new(Member)
		err := rows.Scan(&member.Id, &member.Name)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		members = append(members, member)
	}

	//参加者一覧を返却
	return members
}

func (member *Member) MemberAutoMapperForModel() map[string]interface{} {

	//マップデータの初期化
	data := map[string]interface{}{}

	//各プロパティに対して値がセットされている場合、カラム名に紐づけてマッピング
	if member.Id != 0 {
		data[member_columns[0]] = member.Id
	}

	if member.EventId != 0 {
		data[member_columns[1]] = member.EventId
	}

	if member.Name != "" {
		data[member_columns[2]] = member.Name
	}

	if !member.CreateTime.IsZero() {
		data[member_columns[3]] = member.CreateTime
	}

	if !member.UpdateTime.IsZero() {
		data[member_columns[4]] = member.UpdateTime
	}

	//マッピングデータを返却
	return data

}
