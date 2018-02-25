# Controllers
The purpose of a controller is primary routing the requests around the web app.  The `src\buildtracker\controller\controller.go` file contain the `Startup` func which is called in main.go to start the controller.  

Let's take a look at the code contained in the `controller.go` file and explain what's going on.

```
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
```
First, we configure the variable `homeController` of struct `home` which is defined back in the `...\controller\home.go`.
```
type home struct {
	homeTemplate *template.Template
}
```
The func `Startup` takes a map of templates as a parameter which is the first thing our web application does back at main which is load the templates in the `populateTemplates` func.  

Next, we set the `homeTemplate` element on the `homeController` struct to the template loaded earlier.  Then, we register the routes which each controller individually defines.  For the home controller it is defined as follows.
```
func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}
```
We don't want to forget about some static routes like images and css resources.  

> The paths are important, remember when you execute the build tracker web app it will start in the \bin directory, just take note of the paths.  This was the most challenging part of getting this to work.  
