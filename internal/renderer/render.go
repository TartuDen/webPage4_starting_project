package renderer

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/TartuDen/webPage4_starting_project/internal/config"
	"github.com/TartuDen/webPage4_starting_project/internal/models"
	"github.com/justinas/nosurf"
)

// this var serves to pass data from main.go to render.go
var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RendererTemplate renders template using html/template
func RendererTemplate(w http.ResponseWriter, tmpl string, r *http.Request, td *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		//get the template cache from AppConfig8
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(ok)
	}

	td = AddDefaultData(td, r)

	//optional final error check
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("error writting template to brwoser ", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.html(or tmpl) from ./template
	pages, err := filepath.Glob("./template/*.page.html")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.html (or tmpl)
	for _, page := range pages {
		//here page - is a full path to the file, and we need only name of the file
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./template/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./template/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet

	}
	return myCache, nil

}
