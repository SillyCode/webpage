package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SillyCode/webpage/pkg/config"
	"github.com/SillyCode/webpage/pkg/handlers"
	"github.com/SillyCode/webpage/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Application started on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
