package http

import (
	"fmt"
	"html/template" // Only import if template parsing is done here
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

type PlayController struct{}

// PlayGame renders the play.html template, which embeds the game content in an iframe.
// Templates are parsed on each request, mirroring the original SoftwareController's approach.
func (c *PlayController) PlayGame(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "Game name not provided", http.StatusBadRequest)
		return
	}

	// Parse templates on each request
	tmpl, err := template.ParseFiles("http/views/layouts/main.html", "http/views/play.html")
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

// ServeGameContent serves the static files for a game from the html/$NAME/ directory.
func (c *PlayController) ServeGameContent(w http.ResponseWriter, r *http.Request) {
	gamePath := os.Getenv("GAMES_DIR")
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "Game name not provided", http.StatusBadRequest)
		return
	}

	htmlBaseDir := filepath.Join(gamePath, "html", name)

	// Check if the directory exists and is accessible
	if _, err := os.Stat(htmlBaseDir); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("HTML content for game '%s' not found.", name), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Error accessing HTML content for game '%s': %v", name, err), http.StatusInternalServerError)
		return
	}

	fs := http.StripPrefix(fmt.Sprintf("/play/%s/content", name), http.FileServer(http.Dir(htmlBaseDir)))
	fs.ServeHTTP(w, r)
}
