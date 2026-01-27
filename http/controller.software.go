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

func (c *SoftwareController) index(w http.ResponseWriter, r *http.Request) {
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

type SoftwareShowData struct {
	Software      *domain.Software
	LatestVersion string
}

func (c *SoftwareController) show(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	software, err := c.softwareService.GetByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if software == nil {
		http.NotFound(w, r)
		return
	}

	latest_release, err := c.softwareService.GetLatestRelease(software.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if latest_release == nil {
		http.NotFound(w, r)
		return
	}

	tmpl, _ := template_utils.GetTemplate("software_show", "http/views/shared/layout.html", "http/views/software/show.html")
	tmpl.Execute(w, SoftwareShowData{
		Software:      software,
		LatestVersion: latest_release.Version,
	})
}
