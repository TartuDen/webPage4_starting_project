package main

import (
	"net/http"

	"github.com/TartuDen/webPage4_starting_project/internal/config"
	"github.com/TartuDen/webPage4_starting_project/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(a *config.AppConfig) http.Handler {
	// Create a new instance of the chi router.
	mux := chi.NewRouter()

	middlewareFunc(mux)

	// Add a route for HTTP GET requests to the root path ("/").
	// Associate the MainHandler function from the handler.Repo struct with this route.
	mux.Get("/", handler.Repo.MainHandler)

	// Add a route for HTTP GET requests to the "/about" path.
	// Associate the AboutHandler function from the handler.Repo struct with this route.
	mux.Get("/about", handler.Repo.AboutHandler)

	mux.Get("/dens-room", handler.Repo.DensRoomHandler)

	mux.Get("/yurec-room", handler.Repo.YurecRoomHandler)
	// mux.Post("/yurec-room", handler.Repo.YurecPostRoomHandler)

	mux.Get("/search-availability", handler.Repo.AvailableHandler)

	mux.Get("/bookNow", handler.Repo.BookHandler)
	mux.Post("/bookNow", handler.Repo.BookPostHandler)

	mux.Post("/bookNow-json", handler.Repo.BookJSON)

	mux.Get("/contact", handler.Repo.ContactHandler)

	mux.Get("/make-reservation", handler.Repo.ReservationHandler)


	mux.Post("/make-reservation", handler.Repo.PostMakeReservation)

	mux.Get("/reservation-summary",handler.Repo.ReservationSummary)

	//creating file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Return the configured chi router.
	return mux
}

// middlewareFunc locates all middleware functions
func middlewareFunc(mux *chi.Mux) {
	mux.Use(WriteToConsole)

	// Use the Recoverer middleware from chi.
	// This middleware recovers and logs panics, preventing the application from crashing.
	mux.Use(middleware.Recoverer)

	// Use the NoSurf middleware with the chi router.
	// This adds CSRF protection to all routes registered with this router.
	mux.Use(NoSurf)

	// Apply the SessionLoad middleware to the chi router.
	// This ensures that the loading and saving of user sessions is handled for all routes registered with this router.
	mux.Use(SessionLoad)
}
