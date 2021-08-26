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

	templates := template.Must(template.ParseFiles(tmpl+".gtpl", "view/common/_footer.gtpl", "view/common/_header.gtpl"))

	arr := strings.Split(tmpl, "/")

	file := arr[len(arr)-1]

	templates.ExecuteTemplate(w, "_header.gtpl", nil)
	templates.ExecuteTemplate(w, "_footer.gtpl", nil)
	err := templates.ExecuteTemplate(w, file+".gtpl", i)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//
var validPath = regexp.MustCompile("^/(top|event|member|expense)/([a-zA-Z0-9/]+)$")

//リクエストを受け取りURLに紐づく処理を実行する
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)

		//サーバー起動時にDBコネクションの起動
		model.Connect()
		//処理が完了した際コネクションをクローズする
		defer model.Db.Close()

		if m != nil {
			fn(w, r, m[2])
		} else {
			fn(w, r, "")
		}
	}
}

//エラーハンドラー
func errorHandler(w http.ResponseWriter, path string, status map[interface{}]interface{}, errs map[string]string) {

	if status != nil {
		status["Errors"] = errs
	} else {
		//連携値がない場合マップを初期化する
		status = map[interface{}]interface{}{}
		status["Errors"] = errs
	}
	RenderTemplate(w, path, status)
}
