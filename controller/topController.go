package controller

import "net/http"

func TopHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "top", nil)
}
