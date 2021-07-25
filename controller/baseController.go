package controller

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"text/template"
)

//レスポンスで返すページの構造体定義
type Page struct {
	Title string
	Body  []byte
}

//テンプレートのキャッシュの作成
var templates = template.Must(template.ParseFiles("view/edit.html", "view/view.html"))

//入力値保存関数（代替的にtxtファイルに保存）
/* func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("../"+filename, p.Body, 0600)
} */

//データの読み込み（大体的にテキストファイルの読み込み）
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//テンプレートファイルの読み込み関数
func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	/* t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p) */

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//
var validPath = regexp.MustCompile("^/(edit|save|view|create)/([a-zA-Z0-9]+)$")

//
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
