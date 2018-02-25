# About the Files

## helper.go
I created `helper.go` to help with code readability.  It contains code that is generally called just a few times during the life of the application.  

**populateTemplate()** - We start with the `populateTemplates` func which loads templates and prepares them for the webserver, this is just a map of templates.  This is call in `main.go`.

## main.go
This is the entry point to the application.  Everything starts in the `main` func.  In this application we populate a map of templates, start the controller then listen and serve on port 8000.  I like to keep the amount of code executed from main.go relatively light.    

