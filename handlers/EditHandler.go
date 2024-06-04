package handlers

import (
	"net/http"
	"web-app/entities"
	"web-app/utils"
)

func EditHandler(writer http.ResponseWriter, request *http.Request, title string) {

	page, err := entities.LoadPage(title)

	if err != nil {
		page = &entities.Page{Title: title}
	}

	utils.RenderTemplate("edit", writer, page)
}
