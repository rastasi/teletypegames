package http

import (
	"net/http"

	"teletype_softwares/domain"
	"teletype_softwares/lib/template_utils"
)

type RootController struct {
	softwareService domain.SoftwareServiceInterface
}

func NewRootController(softwareService domain.SoftwareServiceInterface) *RootController {
	return &RootController{
		softwareService: softwareService,
	}
}

func (c *RootController) Index(w http.ResponseWriter, r *http.Request) {
	softwares, err := c.softwareService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Softwares": softwares,
	}

	tmpl, err := template_utils.GetTemplate("root_index", "http/views/shared/layout.html", "http/views/root/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
