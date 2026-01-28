package http

import (
	"net/http"

	"teletype_softwares/domain"
	"teletype_softwares/lib/template_utils"

	"github.com/go-chi/chi/v5"
)

type SoftwareController struct {
	softwareService domain.SoftwareServiceInterface
}

func NewSoftwareController(software_service domain.SoftwareServiceInterface) *SoftwareController {
	return &SoftwareController{softwareService: software_service}
}

func (c *SoftwareController) Index(w http.ResponseWriter, r *http.Request) {
	softwares, err := c.softwareService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template_utils.GetTemplate("software_index", "http/views/shared/layout.html", "http/views/software/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, softwares)
}

func (c *SoftwareController) Show(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	data, err := c.softwareService.GetForShowByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template_utils.GetTemplate("software_show", "http/views/shared/layout.html", "http/views/software/show.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)

}
