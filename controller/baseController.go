package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"text/template"
)

//テンプレートのキャッシュの作成
var templates = template.Must(template.ParseFiles(
	"view/edit.html",
	"view/view.html",
	"view/user/create.html",
	"view/user/complete.html",
	"view/user/index.html"))

//テンプレートファイルの読み込み関数
func RenderTemplate(w http.ResponseWriter, tmpl string, i interface{}) {
	/* t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p) */

	err := templates.ExecuteTemplate(w, tmpl+".html", i)
	if err != nil {
		fmt.Println("error:not found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//
var validPath = regexp.MustCompile("^/(edit|save|view|create|index)/([a-zA-Z0-9]+)$")

//
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			fmt.Println("error:No Path")
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

//入力値保存関数（代替的にtxtファイルに保存）
/* func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("../"+filename, p.Body, 0600)
} */

//データの読み込み（大体的にテキストファイルの読み込み）
// func LoadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }
