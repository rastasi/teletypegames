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

	data := map[string]interface{}{
		"Softwares": softwares,
	}

	tmpl, err := template_utils.GetTemplate("software_index", "http/views/shared/layout.html", "http/views/software/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func (c *SoftwareController) Show(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	showData, err := c.softwareService.GetForShowByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	softwares, err := c.softwareService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Software":           showData.Software,
		"Releases":           showData.Releases,
		"LatestRelease":      showData.LatestRelease,
		"WebPlayableRelease": showData.WebPlayableRelease,
		"Softwares":          softwares,
	}

	tmpl, err := template_utils.GetTemplate("software_show", "http/views/shared/layout.html", "http/views/software/show.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)

}
