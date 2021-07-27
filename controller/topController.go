package controller

import "net/http"

func TopHandler(w http.ResponseWriter, r *http.Request, title string) {
	RenderTemplate(w, "top", nil)
}
