package controller

import (
	"fmt"
	"net/http"
	"shareTravel/common"
	"shareTravel/model"
	"shareTravel/validate"
	"strconv"
	"strings"
)

func MemberHandler(w http.ResponseWriter, r *http.Request, path string) {
	arr := strings.Split(path, SLASH)
	switch arr[0] {
	case ADD:
		memberAddHandler(w, r)
	case SAVE:
		memberSaveHandler(w, r)
	}
}

//参加者追加ハンドラー
func memberAddHandler(w http.ResponseWriter, r *http.Request) {
	event := new(model.Event)
	event.Id, _ = strconv.Atoi(common.GetQueryParam(r))

	status := autoMapperForView(event)

	//参加者追加画面をレンダリング
	RenderTemplate(w, MEMBER_ADD_PATH, status)
}

//参加者保存ハンドラー
func memberSaveHandler(w http.ResponseWriter, r *http.Request) {

	//参加者ポインタ
	member := new(model.Member)

	//入力値を設定
	member.Name = r.FormValue("name")

	//バリデーションチェック
	if ok, errs := validate.MemberValidater(member); !ok {
		//バリデーションエラーがあった場合、エラーハンドリング
		errorHandler(w, MEMBER_ADD_PATH, nil, errs)
		return
	}

	//クエリパラメータの取得、構造体への設定
	str_event_id := common.GetQueryParam(r)
	member.EventId, _ = strconv.Atoi(str_event_id)

	//参加者保存処理
	err := member.SaveMember()

	event := new(model.Event)
	event.Id = member.EventId

	if err != nil {
		errs := map[string]string{}
		errs["Error"] = "予期せぬエラーが発生しました。"
		status := autoMapperForView(event)
		errorHandler(w, ERROR_PATH, status, errs)
		return
	}

	status := autoMapperForView(event, member)

	fmt.Println(status)

	RenderTemplate(w, MEMBER_COMPLETE_PATH, status)
}
