package http

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"teletype_softwares/domain"

	"github.com/go-chi/chi/v5"
)

type DownloadController struct {
	downloadService domain.DownloadService
}

func (c *DownloadController) serveReleaseFile(w http.ResponseWriter, r *http.Request, release *domain.Release, isSource bool) {
	var filePath string
	if isSource {
		filePath = release.SourcePath
	} else {
		filePath = release.CartridgePath
	}

	// Security check: ensure the file is within the softwares directory
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		http.Error(w, "Invalid file path", http.StatusInternalServerError)
		return
	}
	absSoftwaresDir, err := filepath.Abs(os.Getenv("GAMES_DIR"))
	if err != nil {
		http.Error(w, "Invalid softwares directory path", http.StatusInternalServerError)
		return
	}
	if !strings.HasPrefix(absFilePath, absSoftwaresDir) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	http.ServeFile(w, r, filePath)
}

func (c *DownloadController) DownloadSource(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	release, err := c.downloadService.GetLatestRelease(name)
	if err != nil {
		if err == domain.ErrSoftwareNotFound || err == domain.ErrNoReleasesFound {
			http.Error(w, "Software not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.serveReleaseFile(w, r, release, true)
}

func (c *DownloadController) DownloadCartridge(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	release, err := c.downloadService.GetLatestRelease(name)
	if err != nil {
		if err == domain.ErrSoftwareNotFound || err == domain.ErrNoReleasesFound {
			http.Error(w, "Software not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.serveReleaseFile(w, r, release, false)
}

func (c *DownloadController) DownloadSourceByVersion(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	version := chi.URLParam(r, "version")
	release, err := c.downloadService.GetSpecificRelease(name, version)
	if err != nil {
		if err == domain.ErrSoftwareNotFound || err == domain.ErrReleaseNotFound {
			http.Error(w, "Software or release not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.serveReleaseFile(w, r, release, true)
}

func (c *DownloadController) DownloadCartridgeByVersion(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	version := chi.URLParam(r, "version")
	release, err := c.downloadService.GetSpecificRelease(name, version)
	if err != nil {
		if err == domain.ErrSoftwareNotFound || err == domain.ErrReleaseNotFound {
			http.Error(w, "Software or release not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.serveReleaseFile(w, r, release, false)
}
