package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/devillies/bookings/pkg/config"
	"github.com/devillies/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {

	return td

}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {

		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//create template cache

	//get request template from cache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("cant get template")
	}
	buff := new(bytes.Buffer)

	td = addDefaultData(td)
	err := t.Execute(buff, td)

	if err != nil {
		log.Println(err)
	}

	//render the template

	_, err = buff.WriteTo(w)
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
		filename := filepath.Base(page)
		ts, err := template.New(filename).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		var layouts []string
		layouts, err = filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		cache[filename] = ts
	}
	return cache, nil
}
