package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"text/template"
)

//テンプレートのキャッシュの作成
var templates = template.Must(template.ParseFiles(
	"view/top/top.gtpl",
	"view/edit.gtpl",
	"view/view.gtpl",
	"view/user/create.gtpl",
	"view/user/complete.gtpl",
	"view/user/index.gtpl",
	"view/common/_footer.gtpl",
	"view/common/_header.gtpl"))

//テンプレートファイルの読み込み関数
func RenderTemplate(w http.ResponseWriter, tmpl string, i interface{}) {
	/* t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p) */

	templates.ExecuteTemplate(w, "_header.gtpl", i)
	templates.ExecuteTemplate(w, "_footer.gtpl", i)
	err := templates.ExecuteTemplate(w, tmpl+".gtpl", i)
	if err != nil {
		fmt.Println("error:not found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//
var validPath = regexp.MustCompile("^/(top|edit|save|view|create|index)/([a-zA-Z0-9]+)$")

//
func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Println(r.URL.Path)
		if m == nil {
			fmt.Println("error:No Path")
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

//入力値保存関数（代替的にtxtファイルに保存）
/* func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("../"+filename, p.Body, 0600)
} */

//データの読み込み（大体的にテキストファイルの読み込み）
// func LoadPage( ) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }
