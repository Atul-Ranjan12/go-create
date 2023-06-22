package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// Handler function for the CSRF Token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// Funciton to load and save the session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// Checks for authentication
func Auth(next http.Handler) {

}
