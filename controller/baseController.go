package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"shareTravel/model"
	"strings"
	"text/template"
)

//テンプレートファイルの読み込み関数
func RenderTemplate(w http.ResponseWriter, tmpl string, i interface{}) {

	templates := template.Must(template.ParseFiles(tmpl+".gtpl", SIDEBAR_PATH, FOOTER_PATH, HEADER_PATH))

	arr := strings.Split(tmpl, SLASH)

	file := arr[len(arr)-1]

	templates.ExecuteTemplate(w, "_sidebar.gtpl", nil)
	templates.ExecuteTemplate(w, "_header.gtpl", nil)
	templates.ExecuteTemplate(w, "_footer.gtpl", nil)
	err := templates.ExecuteTemplate(w, file+".gtpl", i)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//
var validPath = regexp.MustCompile("^/(event|member|expense|error)/([a-zA-Z0-9/]+)$")

//リクエストを受け取りURLに紐づく処理を実行する
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)

		if m != nil {
			fn(w, r, m[2])
		} else {
			fn(w, r, "")
		}
	}
}

func autoMapperForView(elements ...interface{}) map[string]interface{} {
	status := map[string]interface{}{}

	for _, element := range elements {
		switch element.(type) {
		case *model.Event:
			status["Event"] = element
		case *model.Expense:
			status["Expense"] = element
		case []*model.Expense:
			status["Expenses"] = element
		case *model.Member:
			status["Member"] = element
		case []*model.Member:
			status["Members"] = element
		case *model.MemberExpense:
			status["MemberExpense"] = element
		}
	}

	return status
}

//エラーハンドラー
func errorHandler(w http.ResponseWriter, path string, status map[string]interface{}, errs map[string]string) {

	if status != nil {
		status["Errors"] = errs
	} else {
		//連携値がない場合マップを初期化する
		status = map[string]interface{}{}
		status["Errors"] = errs
	}
	RenderTemplate(w, path, status)
}
