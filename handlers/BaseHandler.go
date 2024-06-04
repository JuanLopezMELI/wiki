package handlers

import (
	"errors"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func GetTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	match := validPath.FindStringSubmatch(r.URL.Path)

	if match == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid page title")
	}

	return match[2], nil
}

func MakeHandler(fn func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := GetTitle(w, r)
		if err != nil {
			return
		}
		fn(w, r, title)
	}
}
