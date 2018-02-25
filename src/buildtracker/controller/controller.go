package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
)

// Startup prepares the controllers to handle and route requests.
func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	homeController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("../public")))
	http.Handle("/css/", http.FileServer(http.Dir("../public")))
}
