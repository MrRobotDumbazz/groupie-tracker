package app

import "net/http"

func Handlers() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/artists/", artist_page)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
}
