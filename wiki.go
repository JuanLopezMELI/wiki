package main

import (
	"net/http"
	"web-app/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/page/", handlers.PageHandler)
	http.HandleFunc("/view/", handlers.MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", handlers.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", handlers.MakeHandler(handlers.SaveHandler))

	http.ListenAndServe(":8080", nil)
}
