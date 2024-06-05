package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // Ensure the cookie is not accessible via JavaScript
		Path:     "/",                  // Set the cookie to be accessible site-wide
		Secure:   app.InProduction,     // Indicate whether the cookie should only be sent over HTTPS
		SameSite: http.SameSiteLaxMode, // Provide some protection against CSRF attacks
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
