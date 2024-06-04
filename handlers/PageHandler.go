package handlers

import "net/http"

func PageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/page/"):]
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
