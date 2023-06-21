package app

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/module"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != "GET" {
		fmt.Println(r.Method)
		Errors(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		Errors(w, r, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Errors(w, r, http.StatusInternalServerError)
		return
	}
	err = module.GetRelation()
	if err != nil {
		Errors(w, r, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, &module.Module)
	if err != nil {
		Errors(w, r, http.StatusInternalServerError)
		return
	}
}
