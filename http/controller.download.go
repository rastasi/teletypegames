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
	service domain.DownloadServiceInterface
}

func NewDownloadController(service domain.DownloadServiceInterface) *DownloadController {
	return &DownloadController{service: service}
}

func (c *DownloadController) serve(w http.ResponseWriter, r *http.Request, release *domain.Release, isSource bool) {
	var filePath string
	if isSource {
		filePath = release.SourcePath
	} else {
		filePath = release.CartridgePath
	}

	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		http.Error(w, "Invalid file path", http.StatusInternalServerError)
		return
	}

	absContentsDir, err := filepath.Abs(os.Getenv("CONTENTS_DIR"))
	if err != nil {
		http.Error(w, "Invalid contents directory path", http.StatusInternalServerError)
		return
	}

	if !strings.HasPrefix(absFilePath, absContentsDir) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	http.ServeFile(w, r, filePath)
}

func (c *DownloadController) handleError(w http.ResponseWriter, err error) {
	if err == domain.ErrSoftwareNotFound || err == domain.ErrNoReleasesFound {
		http.Error(w, "Software not found", http.StatusNotFound)
		return
	}
	if err == domain.ErrReleaseNotFound {
		http.Error(w, "Software or release not found", http.StatusNotFound)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (c *DownloadController) GetLatestSource(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	release, err := c.service.GetLatestRelease(name)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.serve(w, r, release, true)
}

func (c *DownloadController) GetLatestCartridge(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	release, err := c.service.GetLatestRelease(name)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.serve(w, r, release, false)
}

func (c *DownloadController) GetSource(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	version := chi.URLParam(r, "version")
	release, err := c.service.GetSpecificRelease(name, version)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.serve(w, r, release, true)
}

func (c *DownloadController) GetCartridge(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	version := chi.URLParam(r, "version")
	release, err := c.service.GetSpecificRelease(name, version)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.serve(w, r, release, false)
}
