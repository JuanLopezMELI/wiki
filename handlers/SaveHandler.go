package handlers

import (
	"net/http"
	"web-app/entities"
)

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")

	page := &entities.Page{Title: title, Body: []byte(body)}
	err := page.Save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
