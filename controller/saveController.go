package controller

import "net/http"

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// body := r.FormValue("body")

	// err := p.Save()

	/* 	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} */

	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}
