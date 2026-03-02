package http

import (
	"net/http"
	"os"

	"teletype_softwares/domain"
)

type SoftwareUpdaterController struct {
	softwareUpdaterService domain.SoftwareUpdaterServiceInterface
	softwareService        domain.SoftwareServiceInterface
}

func NewSoftwareUpdaterController(
	softwareUpdaterService domain.SoftwareUpdaterServiceInterface,
	softwareService domain.SoftwareServiceInterface,
) *SoftwareUpdaterController {
	return &SoftwareUpdaterController{
		softwareUpdaterService: softwareUpdaterService,
		softwareService:        softwareService,
	}
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

	if err := c.softwareUpdaterService.Update(platform, name, version); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updated"))
}
