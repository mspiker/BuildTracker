# Templates
Templates are simply html files with some extra markup to display data passed as the data context.  In the root of the templates folder you could have the following files.
* **_layout.html** - Establishes the overall look and feel of your site.  It is the shell of the page.  
* **_header.html** - The top of the page, the naviagation bar, login button, search, hamburger menu or logo.  
* **_footer.html** - The bottom of the page.  Repeats on all the pages; socila media links, copyrights, disclaimers, etc.  
* **_sidebar.html** - If you have sidebar navigation you may consider having this as a component.

You will only find _layout and _header in this application.  

## Content Folder
The `content` folder is the page contents.  Here is an example of the `templates\content\home.html` page.
```
{{define "content"}}
    <div class="starter-template">
      <h1>Welcome</h1>
      <p class="lead">Learning the Go programming language is fun.<br> This application uses Bootstrap!</p>
      <a href="shop">Go to the shop</a>
    </div>
{{end}}
```
You will notice our layout is simple, the content section starts with the markup `{{ define "content" }}` and ends with `{{end}}`, you can get more sophisticated with multiple sections.  
