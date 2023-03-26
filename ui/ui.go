package ui

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed "templates"
var templateFS embed.FS

// var templates map[string]*template.Template

// do not put *.tpml or *.html files as templateNames
func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data any) {

	tmpl, err := template.New(templateName).ParseFS(templateFS, "templates/"+templateName+".html")
	if err != nil {
		log.Println(err)
		// return err
	}

	err = tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Println(err)
		// return err
	}
}
