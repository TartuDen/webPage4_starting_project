package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TartuDen/webPage4_starting_project/internal/config"
	"github.com/TartuDen/webPage4_starting_project/internal/handler"
	"github.com/TartuDen/webPage4_starting_project/internal/helpers"
	"github.com/TartuDen/webPage4_starting_project/internal/models"
	"github.com/TartuDen/webPage4_starting_project/internal/renderer"
	"github.com/alexedwards/scs/v2"
)

const Port = ":8080"

// In the main function, an instance of the config.AppConfig struct is created and stored in the app variable.
var app config.AppConfig

var session *scs.SessionManager

var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    Port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	//What are we going to store in session
	gob.Register(models.Reservation{})

	//Change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// The script attempts to create a template cache using renderer.CreateTemplateCache() and stores it in the app.TemplateCache.
	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
		return err
	}
	app.TemplateCache = tc
	//this var to set use cache true or false, when in Dev mode
	app.UseCache = false

	// The handler.NewRepo(&app) call creates a new repository (*Repository)
	// and passes the app configuration to it. This repository is responsible for holding application configuration and handling requests.
	// The NewRepo function in the handler package is used to create a new repository.
	repo := handler.NewRepo(&app)

	// The handler.NewHandlers(repo) function is called to set the global Repo variable to
	// the newly created repository. This allows other parts of the code to access the repository.
	handler.NewHandlers(repo)

	//The renderer.NewTemplate(&app) function is called to set up
	// templates with the app configuration. This prepares the application to render HTML templates.
	renderer.NewTemplate(&app)

	helpers.NewHelpers(&app)

	// http.HandleFunc("/", handler.Repo.MainHandler)
	// http.HandleFunc("/about", handler.Repo.AboutHandler)
	fmt.Println("Server started on Port:", Port)
	// err = http.ListenAndServe(Port, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return nil
}
