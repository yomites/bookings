package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

/**
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}
**/

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	/**
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Just like the name implies, this adds a CSRF token to all POST requests
		csrfHandler := nosurf.New(next)
		csrfHandler.SetBaseCookie(http.Cookie{
			HttpOnly: true,
			Path:     "/",
			Secure:   false, // Set to true in production with HTTPS
			SameSite: http.SameSiteLaxMode,
		})
		csrfHandler.ServeHTTP(w, r)
	})
	**/ // I have refactored the code above to remove the unnecessary wrapping of the http.HandlerFunc and to directly return the csrfHandler. This simplifies the code and achieves the same functionality.
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
