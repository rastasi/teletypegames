package http

import (
	"net/http"
	"os"

	"teletype_softwares/domain"
)

type SoftwareUpdaterController struct {
	service domain.SoftwareUpdaterServiceInterface
}

func NewSoftwareUpdaterController(service domain.SoftwareUpdaterServiceInterface) *SoftwareUpdaterController {
	return &SoftwareUpdaterController{service: service}
}

func (c *SoftwareUpdaterController) Update(w http.ResponseWriter, r *http.Request) {
	secret := r.URL.Query().Get("secret")
	if secret != os.Getenv("UPDATE_SECRET") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	platform := r.URL.Query().Get("platform")
	name := r.URL.Query().Get("name")
	version := r.URL.Query().Get("version")

	if version == "" {
		http.Error(w, "Version not provided", http.StatusBadRequest)
		return
	}

	if err := c.service.Update(platform, name, version); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updated"))
}
