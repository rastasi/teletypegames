package http

import (
	"html/template"
	"net/http"

	"teletype_softwares/domain"

	"github.com/go-chi/chi/v5"
)

type SoftwareController struct {
	softwareService domain.SoftwareService
}

func (c *SoftwareController) index(w http.ResponseWriter, r *http.Request) {
	softwares, err := c.softwareService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("http/views/layouts/main.html", "http/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, softwares)
}

func (c *SoftwareController) releases(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	software, err := c.softwareService.GetByNameWithReleases(name)
	if err != nil {
		if err == domain.ErrSoftwareNotFound {
			http.Error(w, "Software not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("http/views/layouts/main.html", "http/views/releases.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
		"Software": software,
	})
}
