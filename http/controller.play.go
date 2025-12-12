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

func (c *PlayController) Play(w http.ResponseWriter, r *http.Request) {
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

	tmpl, err := template_utils.GetTemplate("play_controller_play", "http/views/shared/layout.html", "http/views/play/play.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
		"Software": software,
	})
}

func (c *PlayController) ServeContent(w http.ResponseWriter, r *http.Request) {
	contentsPath := os.Getenv("GAMES_DIR")
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}

	htmlBaseDir := filepath.Join(contentsPath, "html", name)

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
