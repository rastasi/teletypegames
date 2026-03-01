package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"teletype_softwares/domain"
)

type APISoftwareController struct {
	softwareService domain.SoftwareServiceInterface
}

func NewAPISoftwareController(software_service domain.SoftwareServiceInterface) *APISoftwareController {
	return &APISoftwareController{softwareService: software_service}
}

func (c *APISoftwareController) Index(w http.ResponseWriter, r *http.Request) {
	softwares, _ := c.softwareService.DetailedList()

	for _, softwareShowData := range softwares.Softwares {
		for i := range softwareShowData.Releases {
			release := &softwareShowData.Releases[i]
			release.CartridgePath = strings.ReplaceAll(release.CartridgePath, "/softwares/", "/file/")
			release.SourcePath = strings.ReplaceAll(release.SourcePath, "/softwares/", "/file/")
			release.HTMLFolderPath = strings.ReplaceAll(release.HTMLFolderPath, "/softwares/", "/file/")
			release.DocsFolderPath = strings.ReplaceAll(release.DocsFolderPath, "/softwares/", "/file/")
		}
		if softwareShowData.LatestRelease != nil {
			release := softwareShowData.LatestRelease
			release.CartridgePath = strings.ReplaceAll(release.CartridgePath, "/softwares/", "/file/")
			release.SourcePath = strings.ReplaceAll(release.SourcePath, "/softwares/", "/file/")
			release.HTMLFolderPath = strings.ReplaceAll(release.HTMLFolderPath, "/softwares/", "/file/")
			release.DocsFolderPath = strings.ReplaceAll(release.DocsFolderPath, "/softwares/", "/file/")
		}
		if softwareShowData.WebPlayableRelease != nil {
			release := softwareShowData.WebPlayableRelease
			release.CartridgePath = strings.ReplaceAll(release.CartridgePath, "/softwares/", "/file/")
			release.SourcePath = strings.ReplaceAll(release.SourcePath, "/softwares/", "/file/")
			release.HTMLFolderPath = strings.ReplaceAll(release.HTMLFolderPath, "/softwares/", "/file/")
			release.DocsFolderPath = strings.ReplaceAll(release.DocsFolderPath, "/softwares/", "/file/")
		}
	}

	json.NewEncoder(w).Encode(softwares)
}
