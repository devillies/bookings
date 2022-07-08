package handlers

import (
	"net/http"

	"github.com/devillies/bookings/pkg/config"
	"github.com/devillies/bookings/pkg/models"
	"github.com/devillies/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// Template Data

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIp := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := repo.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := map[string]string{
		"test":      "hello,again",
		"remote_ip": remoteIp,
	}

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
