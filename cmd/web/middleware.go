package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole is a middleware that logs a message to the console
// each time an HTTP request is received. It takes an `http.Handler` as
// an argument, allowing it to be chained with other middleware or handlers.
func WriteToConsole(next http.Handler) http.Handler {
	// Return an http.Handler that logs a message to the console and then
	// calls the ServeHTTP method of the next handler in the chain.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log a message to the console indicating that the page has been hit.
		// fmt.Println("middleware is working!")

		// Call the ServeHTTP method of the next handler in the chain
		// to continue processing the HTTP request.
		next.ServeHTTP(w, r)
	})
}

// NoSurf is a middleware function that adds CSRF protection to the provided http.Handler.
func NoSurf(next http.Handler) http.Handler {
	// Create a new nosurf middleware, wrapping the provided http.Handler.
	csrfHandler := nosurf.New(next)

	// Set the properties of the base CSRF cookie.
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // The cookie is not accessible through JavaScript.
		Path:     "/",                  // The cookie is available for all paths on the domain.
		Secure:   app.InProduction,     // Change to true if serving over HTTPS.
		SameSite: http.SameSiteLaxMode, // SameSiteLaxMode is a common setting for CSRF protection.
	})

	// Return the configured CSRF handler.
	return csrfHandler
}

// SessionLoad is a middleware function that manages the loading and saving of user sessions.
// It takes an http.Handler as an argument and returns a new http.Handler.
func SessionLoad(next http.Handler) http.Handler {
	// The session.LoadAndSave function is used to handle session loading and saving.
	// It is likely provided by a session management library.
	return session.LoadAndSave(next)
}
