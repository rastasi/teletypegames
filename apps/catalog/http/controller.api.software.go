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

func replaceReleasePaths(release *domain.Release) {
	release.CartridgePath = strings.ReplaceAll(release.CartridgePath, "/softwares/", "/file/")
	release.SourcePath = strings.ReplaceAll(release.SourcePath, "/softwares/", "/file/")
	release.HTMLFolderPath = strings.ReplaceAll(release.HTMLFolderPath, "/softwares/", "/file/")
	release.DocsFolderPath = strings.ReplaceAll(release.DocsFolderPath, "/softwares/", "/file/")
}

func (c *APISoftwareController) Index(w http.ResponseWriter, r *http.Request) {
	softwares, _ := c.softwareService.DetailedList()

	for idx := range softwares.Softwares {
		softwareShowData := &softwares.Softwares[idx]

		for i := range softwareShowData.Releases {
			replaceReleasePaths(&softwareShowData.Releases[i])
		}

		// LatestRelease és WebPlayableRelease a Releases slice elemeire mutat,
		// ezért ezeket külön is frissíteni kell, mivel értékmásolat kerülhet bele
		if softwareShowData.LatestRelease != nil {
			replaceReleasePaths(softwareShowData.LatestRelease)
		}
		if softwareShowData.WebPlayableRelease != nil {
			replaceReleasePaths(softwareShowData.WebPlayableRelease)
		}
	}

	json.NewEncoder(w).Encode(softwares)
}
