package http

import (
	"fmt"
	"net/http"
	"os"
	"teletype_softwares/domain"
	"teletype_softwares/lib/template_utils"

	"github.com/go-chi/chi/v5"
)

type PlayController struct {
	softwareService domain.SoftwareServiceInterface
	fileService     domain.FileServiceInterface
}

func NewPlayController(softwareService domain.SoftwareServiceInterface, fileService domain.FileServiceInterface) *PlayController {
	return &PlayController{
		softwareService: softwareService,
		fileService:     fileService,
	}
}

func (c *PlayController) Play(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	version := chi.URLParam(r, "version")

	software, err := c.softwareService.GetByNameWithReleases(name)
	if err != nil {
		if err == domain.ErrSoftwareNotFound {
			http.Error(w, "Software not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var webPlayableRelease *domain.Release
	for _, release := range software.Releases {
		if release.Version == version && release.HTMLFolderPath != "" {
			webPlayableRelease = &release
			break
		}
	}

	if webPlayableRelease == nil {
		http.Error(w, "No web-playable version found for this software", http.StatusNotFound)
		return
	}

	tmpl, err := template_utils.GetTemplate("play_controller_play", "http/views/shared/layout.html", "http/views/play/play.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
		"Software":           software,
		"WebPlayableRelease": webPlayableRelease,
	})
}

func (c *PlayController) ServeContent(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	version := chi.URLParam(r, "version")

	html_base_dir := c.fileService.GetPath(name + "-" + version)

	if _, err := os.Stat(html_base_dir); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("Content for '%s' not found.", name), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Error accessing content for '%s': %v", name, err), http.StatusInternalServerError)
		return
	}

	fs := http.StripPrefix(fmt.Sprintf("/play/%s/%s/content", name, version), http.FileServer(http.Dir(html_base_dir)))
	fs.ServeHTTP(w, r)
}
