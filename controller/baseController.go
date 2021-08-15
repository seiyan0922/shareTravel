package controller

import (
	"fmt"
	"net/http"
	"regexp"
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

//
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
