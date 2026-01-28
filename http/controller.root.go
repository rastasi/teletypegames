package http

import (
	"net/http"

	"teletype_softwares/domain"
	"teletype_softwares/lib/template_utils"
)

type RootController struct {
	service domain.SoftwareServiceInterface
}

func NewRootController() *RootController {
	return &RootController{}
}

func (c *RootController) Index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template_utils.GetTemplate("root_index", "http/views/shared/layout.html", "http/views/root/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
