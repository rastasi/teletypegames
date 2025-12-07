package http

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"teletype_softwares/lib/template_utils"

	"github.com/go-chi/chi/v5"
)

type PlayController struct{}

func NewPlayController() *PlayController {
	return &PlayController{}
}

func (c *PlayController) Play(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}

	tmpl, err := template_utils.GetTemplate("play", "http/views/layouts/main.html", "http/views/play.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Name string
	}{
		Name: name,
	}

	tmpl.Execute(w, data)
}

func (c *PlayController) ServeContent(w http.ResponseWriter, r *http.Request) {
	contentsPath := os.Getenv("CONTENTS_DIR")
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
