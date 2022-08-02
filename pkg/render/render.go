package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sanmitM312/room-booking-app/pkg/config"
	"github.com/sanmitM312/room-booking-app/pkg/models"

)

var app *config.AppConfig

func AddDefaultData(td *models.TemplateData) *models.TemplateData{

	return td
}
// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a 
}
// This renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData){
	var tc map[string]*template.Template

	if app.UseCache{
		// when pulling a page we refer to this application template cache
		tc = app.TemplateCache
	}else{
		tc, _ = CreateTemplateCache()
	}
	

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	//for fine grained error checking 
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	
	// execute the got template in the buf
	_ = t.Execute(buf,td)
	// if err != nil {
	// 	log.Println(err) // indicates error comes from map
	// }
	// render the template

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser",err)
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl,"./templates/base.layout.tmpl" )
	// err := parsedTemplate.Execute(w,nil)

	// if err != nil {
	// 	fmt.Println("error parsing template: ", err)
	// }
}

// complex template Cache starts
func CreateTemplateCache()(map[string]*template.Template,error) {

	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages,err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache,err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages{
		name := filepath.Base(page)// just the filename[last element of the path]
		// store the template the page in template object name
		ts, err := template.New(name).ParseFiles(page)
		// if error happens return
		if err != nil {
			return myCache, err
		}

		//look at the layouts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts,err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache,err
			}
		}

		myCache[name] = ts
	}

	return myCache,nil
}
// complex template Cache ends


// as ParseFiles returns a *template.Template pointer
// SIMPLE TEMPLATE CACHING STARTS
var tc1 = make(map[string]*template.Template)
func createTemplateCacheTest(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s",t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache
	tc1[t] = tmpl

	return nil
}
// simple template cache
func RenderTemplateTest1(w http.ResponseWriter, t string){
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc1[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCacheTest(t)
		if err != nil {
			log.Println(err)
		}
	}else{
		// we have template in cache
		log.Println("using cached template")
	}
	tmpl = tc1[t]
	err = tmpl.Execute(w,nil)
	if err != nil {
		log.Println(err)
	}
}
// SIMPLE TEMPLATE CACHING ENDS

