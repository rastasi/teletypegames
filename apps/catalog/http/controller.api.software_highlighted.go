package http

import (
	"encoding/json"
	"net/http"

	"teletype_softwares/domain"
)

type APISoftwareHighlightedController struct {
	softwareService domain.SoftwareServiceInterface
}

func NewAPISoftwareHighlightedController(software_service domain.SoftwareServiceInterface) *APISoftwareHighlightedController {
	return &APISoftwareHighlightedController{softwareService: software_service}
}

func (c *APISoftwareHighlightedController) Index(w http.ResponseWriter, r *http.Request) {
	softwareShowData, err := c.softwareService.GetHighlighted()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if softwareShowData == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "no highlighted software found"})
		return
	}

	for i := range softwareShowData.Releases {
		replaceReleasePaths(&softwareShowData.Releases[i])
	}

	if softwareShowData.LatestRelease != nil {
		replaceReleasePaths(softwareShowData.LatestRelease)
	}
	if softwareShowData.WebPlayableRelease != nil {
		replaceReleasePaths(softwareShowData.WebPlayableRelease)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(softwareShowData)
}
