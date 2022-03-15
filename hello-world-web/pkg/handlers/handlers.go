package handlers

import (
	"net/http"

	"github.com/ozan1338/study-go/pkg/config"
	"github.com/ozan1338/study-go/pkg/models"
	"github.com/ozan1338/study-go/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the Home Page Handle
func (m *Repository ) Home(w http.ResponseWriter, r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the About Page Handle
func (m *Repository) About(w http.ResponseWriter, r *http.Request){

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

