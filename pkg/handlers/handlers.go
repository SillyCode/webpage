package handlers

import (
	"net/http"

	"github.com/SillyCode/webpage/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template:", err)
// 		return
// 	}
// }
