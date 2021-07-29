package controller

import "net/http"

func TopHandler(w http.ResponseWriter, r *http.Request, path string) {
	RenderTemplate(w, "view/top/top", nil)
}
