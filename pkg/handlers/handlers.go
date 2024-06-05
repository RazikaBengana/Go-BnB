package handlers

import (
	"github.com/RazikaBengana/Go-BnB/pkg/config"
	"github.com/RazikaBengana/Go-BnB/pkg/models"
	"github.com/RazikaBengana/Go-BnB/pkg/render"
	"net/http"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type that holds the application configuration
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository with the given application configuration
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the global repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home handles the home page requests
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About handles the about page requests
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap["remote_ip"] = remoteIP

	// Send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
