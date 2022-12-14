package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/sanmitM312/room-booking-app/internal/helpers"
)

func WritetToConsole(next http.Handler) http.Handler {
	// return an anonymous function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page.")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protetion to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	// log.Fatal("Idhar tak aya")
	// if session != nil{
	// 	fmt.Println(session)
	// }
	return session.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticated(r) {
			session.Put(r.Context(), "error", "Log in first!")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		}
		next.ServeHTTP(w,r)
	})
}
