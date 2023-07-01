package handlers

import (
	"github.com/MelihEmreGuler/web-content-management/pkg/config"
	"github.com/MelihEmreGuler/web-content-management/pkg/models"
	"github.com/MelihEmreGuler/web-content-management/pkg/render"
	"net/http"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the type for the repository pattern
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository for the handlers
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a, // assign the app config to the repository
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r // assign the repository to the handlers
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the about page")
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
