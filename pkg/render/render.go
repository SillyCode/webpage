package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)
	err = template.Execute(buffer, nil)
	if err != nil {
		log.Println(err)
	}

	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templates, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			templates, err = templates.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = templates
	}

	return cache, nil
}
