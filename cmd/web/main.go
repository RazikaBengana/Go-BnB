package main

import (
	"encoding/gob"
	"fmt"
	"github.com/RazikaBengana/Go-BnB/internal/config"
	"github.com/RazikaBengana/Go-BnB/internal/handlers"
	"github.com/RazikaBengana/Go-BnB/internal/models"
	"github.com/RazikaBengana/Go-BnB/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the entry point for the application
func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

// run initializes the application configuration and dependencies
func run() error {
	// What I am going to put in the session
	gob.Register(models.Reservation{})

	// Change this to true when in production to enforce secure settings
	app.InProduction = false

	// Initialize the session manager with a 24-hour session lifetime
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// Store the session manager in the app config
	app.Session = session

	// Create a template cache for rendering HTML templates
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	return nil
}
