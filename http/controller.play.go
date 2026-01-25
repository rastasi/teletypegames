package http

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"teletype_softwares/domain"
	"teletype_softwares/lib/template_utils"

	"github.com/go-chi/chi/v5"
)

type PlayController struct {
	softwareService domain.SoftwareServiceInterface
}

func NewPlayController(softwareService domain.SoftwareServiceInterface) *PlayController {
	return &PlayController{
		softwareService: softwareService,
	}
}

func (c *PlayController) PlayV1(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}

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
		if release.WebPlayable {
			webPlayableRelease = &release
			break
		}
	}

	if webPlayableRelease == nil {
		http.Error(w, "No web-playable version found for this software", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/play/%s/%s", name, webPlayableRelease.Version), http.StatusFound)
}

func (c *PlayController) Play(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}

	version := chi.URLParam(r, "version")
	if version == "" {
		http.Error(w, "Version not provided", http.StatusBadRequest)
		return
	}

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
		if release.Version == version && release.WebPlayable {
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

func (c *PlayController) ServeContentV1(w http.ResponseWriter, r *http.Request) {
	contentsPath := os.Getenv("GAMES_DIR")
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}

	software, err := c.softwareService.GetByNameWithReleases(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Content for '%s' not found.", name), http.StatusNotFound)
		return
	}

	var webPlayableRelease *domain.Release
	for _, release := range software.Releases {
		if release.WebPlayable {
			webPlayableRelease = &release
			break
		}
	}

	if webPlayableRelease == nil {
		http.Error(w, "No web-playable version found for this software", http.StatusNotFound)
		return
	}

	htmlBaseDir := filepath.Join(contentsPath, "html", name, webPlayableRelease.Version)

	if _, err := os.Stat(htmlBaseDir); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("Content for '%s' not found.", name), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Error accessing content for '%s': %v", name, err), http.StatusInternalServerError)
		return
	}

	fs := http.StripPrefix(fmt.Sprintf("/play/%s/content", name), http.FileServer(http.Dir(htmlBaseDir)))
	fs.ServeHTTP(w, r)
}

func (c *PlayController) ServeContent(w http.ResponseWriter, r *http.Request) {
	contentsPath := os.Getenv("GAMES_DIR")
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}

	version := chi.URLParam(r, "version")
	if version == "" {
		http.Error(w, "Version not provided", http.StatusBadRequest)
		return
	}

	htmlBaseDir := filepath.Join(contentsPath, name, fmt.Sprintf("%s-%s-html", name, version))

	if _, err := os.Stat(htmlBaseDir); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("Content for '%s' not found.", name), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Error accessing content for '%s': %v", name, err), http.StatusInternalServerError)
		return
	}

	fs := http.StripPrefix(fmt.Sprintf("/play/%s/%s/content", name, version), http.FileServer(http.Dir(htmlBaseDir)))
	fs.ServeHTTP(w, r)
}
