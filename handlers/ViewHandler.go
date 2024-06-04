package handlers

import (
	"net/http"
	"web-app/entities"
	"web-app/utils"
)

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := entities.LoadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	utils.RenderTemplate("view", w, page)
}
