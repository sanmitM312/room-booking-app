package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sanmitM312/room-booking-app/pkg/config"
	"github.com/sanmitM312/room-booking-app/pkg/handlers"
	"github.com/sanmitM312/room-booking-app/pkg/render"
)


const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
// to run all the go files
// do go run *.go
// go run cmd/web/*.go
func main(){

	// change this to true in production 
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // in production it is true,only allowed https connections 

	app.Session = session
	// as soon as application starts, create the template cache
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("failed to create template cache")
	}

	app.TemplateCache = tc 
	app.UseCache = false
	
	Repo := handlers.NewRepo(&app)
	handlers.NewHandlers(Repo)
	
	// after that render the templates as reqd
	render.NewTemplates(&app)
	
	fmt.Printf("Starting application on port %v", portNumber)
	// start a web server and listen to requests 
	// _ = http.ListenAndServe(portNumber,nil)

	srv := &http.Server {
		Addr : portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}