package render

import (
	"bytes"
	"fmt"
	"github.com/MelihEmreGuler/web-content-management/pkg/config"
	"github.com/MelihEmreGuler/web-content-management/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

var functions = template.FuncMap{}

var app *config.AppConfig //pointer to the app config

// NewTemplates sets the config for the template package
func NewTemplates(config *config.AppConfig) {
	app = config
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	//add default data to the template data
	td.IntMap = make(map[string]int)
	td.IntMap["year"] = time.Now().Year()
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	} else {
		//create a new template cache every time for development
		err := error(nil)
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}

	//get requested page from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer) //create a new buffer

	td = AddDefaultData(td) //set default data

	err := t.Execute(buf, td) //execute the template and write it to the buffer
	if err != nil {
		log.Fatal(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tmpl, "rendered successfully")
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	ceche := map[string]*template.Template{}

	// get all the files named *.page.html from the templates directory
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return ceche, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		// get the base name of the file
		name := filepath.Base(page)
		fmt.Println("page is currently", page)

		// parse the file
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return ceche, err
		}

		// get all the files named *.layout.html from the templates directory
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return ceche, err
		}

		// if there are any files ending with *.layout.html in the templates directory
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return ceche, err
			}
		}

		// add the template to the cache
		ceche[name] = ts
	}

	// return the map and nil error if everything is ok
	return ceche, nil
}

// inefficient way of doing it
/*
var tc = make(map[string]*template.Template) //template cache
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	//check to see if we already have the template in the cache
	_, inMap := tc[t] //returns the template and a boolean (inMap) if it exists in the map
	if !inMap {
		//need to load the template from disk into the cache
		fmt.Println("loading template from disk into cache")
		err = createTemplateCache(t)
	} else {
		//get the template from the cache
		fmt.Println("using template from cache")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	//parse the template files
	tmpl, err := template.ParseFiles(templates...) //... means pass in a slice of strings instead of a single string
	if err != nil {
		return err
	}

	//add the template to the cache
	tc[t] = tmpl
	return nil
}
*/
