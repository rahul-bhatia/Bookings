package handlers

import (
	"fmt"
	"net/http"

	"github.com/rahulb/bookings/pkg/config"
	"github.com/rahulb/bookings/pkg/models"
	"github.com/rahulb/bookings/pkg/render"
)

//TemplateData holds data sent from handlers to template.

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println("Ip is", remoteIP)
	m.App.Session.Put(r.Context(), "RemoteIP", remoteIP)

	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{})
	//fmt.Fprintf(rw, "hello world")
}

func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	//fmt.Println("GET:about")
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again!"

	remoteIp := m.App.Session.GetString(r.Context(), "RemoteIP")
	println("Setting ip from about", remoteIp)
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//this is comment
