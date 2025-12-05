package http

import (
	"net/http"
	"os"

	"teletype_softwares/domain"
)

type SoftwareUpdaterController struct {
	softwareUpdaterService domain.SoftwareUpdaterService
}

func (c *SoftwareUpdaterController) update(w http.ResponseWriter, r *http.Request) {
	secret := r.URL.Query().Get("secret")
	if secret != os.Getenv("UPDATE_SECRET") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	platform := r.URL.Query().Get("platform")
	name := r.URL.Query().Get("name")

	if err := c.softwareUpdaterService.UpdateSoftware(platform, name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updated"))
}
