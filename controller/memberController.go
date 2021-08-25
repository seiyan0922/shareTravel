package controller

import (
	"net/http"
	"shareTravel/model"
	"strconv"
	"strings"
)

func MemberHandler(w http.ResponseWriter, r *http.Request, path string) {
	arr := strings.Split(path, "/")
	switch arr[0] {
	case "add":
		memberAddHandler(w, r)
	case "save":
		memberSaveHandler(w, r)
	}
}

func memberAddHandler(w http.ResponseWriter, r *http.Request) {
	event := new(model.Event)
	query := r.URL.Query().Encode()
	strid := strings.Split(query, "=")[1]
	event.Id, _ = strconv.Atoi(strid)

	RenderTemplate(w, "view/member/add", event)
}

func memberSaveHandler(w http.ResponseWriter, r *http.Request) {
	member := new(model.Member)
	member.Name = r.FormValue("name")

	query := r.URL.Query().Encode()
	strid := strings.Split(query, "=")[1]
	member.EventId, _ = strconv.Atoi(strid)

	member.SaveMember()

	RenderTemplate(w, "view/member/complete", member)
}

func postMembersCnv(str_members string) []model.Member {

	replaced1 := strings.Replace(str_members, "[", "", -1)
	replaced2 := strings.Replace(replaced1, "]", "", -1)
	replaced3 := strings.Replace(replaced2, "{", "", -1)

	members_arr := strings.Split(replaced3, "} ")

	var members []model.Member

	for _, str_member := range members_arr {
		member_arr := strings.Split(str_member, " ")

		member_id, _ := strconv.Atoi(member_arr[0])
		event_id, _ := strconv.Atoi(member_arr[1])
		name := member_arr[2]

		var member model.Member
		member.Id = member_id
		member.EventId = event_id
		member.Name = name

		members = append(members, member)
	}

	return members

}

//参加メンバー負担金総額取得処理
func GetMembersTotal(members []model.Member) []model.Member {
	for i := 0; i < len(members); i++ {
		model.GetMemberExpense(&members[i])
	}

	return members
}
