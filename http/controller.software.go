package http

import (
	"net/http"

	"teletype_softwares/domain"
	"teletype_softwares/lib/template_utils"

	"github.com/go-chi/chi/v5"
)

type SoftwareController struct {
	service domain.SoftwareServiceInterface
}

func NewSoftwareController(service domain.SoftwareServiceInterface) *SoftwareController {
	return &SoftwareController{service: service}
}

func (c *SoftwareController) index(w http.ResponseWriter, r *http.Request) {
	softwares, err := c.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template_utils.GetTemplate("index", "http/views/layouts/main.html", "http/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, softwares)
}

func (c *SoftwareController) releases(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	software, err := c.service.GetByNameWithReleases(name)
	if err != nil {
		if err == domain.ErrSoftwareNotFound {
			http.Error(w, "Software not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template_utils.GetTemplate("releases", "http/views/layouts/main.html", "http/views/releases.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
		"Software": software,
	})
}
