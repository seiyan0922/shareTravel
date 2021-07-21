package controller

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("view/edit.html", "view/view.html"))

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("../"+filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	/* t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p) */

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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
