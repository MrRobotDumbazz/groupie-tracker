package app

import (
	"html/template"
	"net/http"
	"strconv"
)

type Error struct {
	Statusint  int
	StatusText string
	StatusTextandInt string
}

func Errors(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	t, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w, strconv.Itoa(status)+" "+http.StatusText(status), status)
		return
	}
	error1 := Error{status, http.StatusText(status), strconv.Itoa(status) + " " + http.StatusText(status)}
	t.Execute(w, error1)
}
