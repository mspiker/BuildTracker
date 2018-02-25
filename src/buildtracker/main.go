package main

import (
	"firstwebapp/controller"
	"net/http"
)

func main() {
	templates := populateTemplates()
	controller.Startup(templates)
	http.ListenAndServe(":8000", nil)
}
