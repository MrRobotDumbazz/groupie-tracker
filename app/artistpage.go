package app

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/module"
)

func artist_page(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "GET" {
		fmt.Println(r.Method)
		tmpl, err := template.ParseFiles("./templates/artistpage.html")
		if err != nil {
			Errors(w, r, http.StatusInternalServerError)
			return
		}
		id := strings.TrimPrefix(r.URL.Path, "/artists/")
		req, err := http.NewRequest("HEAD", r.URL.Path, nil)
		if err != nil {
			Errors(w, r, http.StatusInternalServerError)
		}
	
		http.DefaultClient.Do(req)
		if id == "" {
			Errors(w, r, http.StatusNotFound)
			return
		}
		artistid, err := strconv.Atoi(id)
		if err != nil {
			Errors(w, r, http.StatusInternalServerError)
			return
		}

		if artistid < 1 || artistid > len(module.Module) {
			Errors(w, r, http.StatusNotFound)
			return
		}
		tmpl.Execute(w, module.Module[artistid-1])
	} else {
		fmt.Println(r.Method)
		Errors(w, r, http.StatusMethodNotAllowed)
		return
	}
}
