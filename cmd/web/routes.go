package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sanmitM312/room-booking-app/pkg/config"
	"github.com/sanmitM312/room-booking-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler{
	//  mux := pat.New() // pat package for rute

	//  mux.Get("/",http.HandlerFunc(handlers.Repo.Home))
	//  mux.Get("/about",http.HandlerFunc(handlers.Repo.About))

	/*------USING CHI ------ */
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WritetToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))
	
	 return mux
}