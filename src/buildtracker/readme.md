# About the Files

## helper.go
I created `helper.go` to help with code readability.  It contains code that is generally called just a few times during the life of the application.  

**populateTemplate()** - We start with the `populateTemplates` func which loads templates and prepares them for the webserver, this is just a map of templates.  This is call in `main.go`.

## main.go
This is the entry point to the application.  Everything starts in the `main` func.  In this application we populate a map of templates, start the controller then listen and serve on port 8000.  I like to keep the amount of code executed from main.go relatively light.  
```
func main() {
	templates := populateTemplates()
	controller.Startup(templates)
	http.ListenAndServe(":8000", nil)
}
```
Our application uses an MVC framework. 
* Models are stored in the `src\buildtracker\model` folder, primarily file containing structs defining the data, we then pass those structs as the data context to the views, or templates.  
* View are the templates.  Templates are simply html files using template fields.  The html files are stored in the `templates` folder.  
* Controllers are stored in the `src\buildtracker\controllers` folder and those act to route requests.  While routing engines like Gorilla/Mux and Http Router could be used I tried to stay vanilla with Go. 

## Controllers and Views
Controllers will route users throughout the application.  Before we can start the web server we first must build a controller to take care of the home page and build the templates, or html pages along with the supporting files such as and images, client-side javascript and style sheets.  Those are saved in the `templates` folder.  See [Templates](../readme.md) for more explination. 

