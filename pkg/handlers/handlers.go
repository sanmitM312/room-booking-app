package handlers

import (
	"fmt"
	"net/http"

	"github.com/sanmitM312/room-booking-app/pkg/config"
	"github.com/sanmitM312/room-booking-app/pkg/models"
	"github.com/sanmitM312/room-booking-app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repositry is the repository type
type Repository struct{
	App *config.AppConfig
}

// Creates a new repository,returns a pointer to the repository
func NewRepo(a *config.AppConfig) *Repository{
	// initialise a struct a return the repository
	return &Repository{
		App : a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}


// never executes unless u do a listen and serve call
func (m *Repository)Home(w http.ResponseWriter, r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)
	
	render.RenderTemplate(w,"home.page.tmpl",&models.TemplateData{})
}

// About is the a about page handler
func (m *Repository)About(w http.ResponseWriter, r *http.Request){
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = remoteIP
	fmt.Println("ha bhai in about hu",remoteIP)

	// send the data to the template of type Template Data
	render.RenderTemplate(w,"about.page.tmpl",&models.TemplateData{
		StringMap: stringMap,
	})
}


