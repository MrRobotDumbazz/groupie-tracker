package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/app"
)

func main() {
	fmt.Println("Server is started localhost:8180")
	app.Handlers()
	err := http.ListenAndServe(":8180", nil)
	if err != nil {
		log.Print(err)
		return
	}
}
