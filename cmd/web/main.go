package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sanmitM312/room-booking-app/internal/config"
	"github.com/sanmitM312/room-booking-app/internal/driver"
	"github.com/sanmitM312/room-booking-app/internal/handlers"
	"github.com/sanmitM312/room-booking-app/internal/helpers"
	"github.com/sanmitM312/room-booking-app/internal/models"
	"github.com/sanmitM312/room-booking-app/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {

	db,err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	defer close(app.MailChan)

	// called after closing the listenig to mailchan channel
	fmt.Println("Starting mail listener...")
	listenForMail()

	// msg := models.MailData{
	// 	To: "john@do.ca",
	// 	From: "me@here.com",
	// 	Subject: "Some subject",
	// 	Content: "",
	// }

	// app.MailChan <- msg // send the mail

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB,error) {
	// things to put inside the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})

	// listens for MailData type 
	mailChan := make(chan models.MailData)
	app.MailChan = mailChan
	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=root")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	
	log.Println("Connected to the database.")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}


	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app,db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	  
	return db,nil
}
