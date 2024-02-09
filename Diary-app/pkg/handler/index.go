package handler

import (
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("../pkg/templates/index.html", "../pkg/templates/show.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"Title": "index"}
	RenderTemplate(w, "index", data)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", data); err != nil {
		log.Fatalln("Unable to execute template.")
	}
}
