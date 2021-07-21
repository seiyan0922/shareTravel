package controller

import "net/http"

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	RenderTemplate(w, "view", p)

}
