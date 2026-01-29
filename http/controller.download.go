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
	downloadService domain.DownloadServiceInterface
	softwareService domain.SoftwareServiceInterface
}

func NewDownloadController(
	downloadService domain.DownloadServiceInterface,
	softwareService domain.SoftwareServiceInterface,
) *DownloadController {
	return &DownloadController{
		downloadService: downloadService,
		softwareService: softwareService,
	}
}

func (c *DownloadController) serve(w http.ResponseWriter, r *http.Request, release *domain.Release, is_source bool) {
	var file_path string
	if is_source {
		file_path = release.SourcePath
	} else {
		file_path = release.CartridgePath
	}

	abs_file_path, err := filepath.Abs(file_path)
	if err != nil {
		http.Error(w, "Invalid file path", http.StatusInternalServerError)
		return
	}

	abs_contents_dir, err := filepath.Abs(os.Getenv("FILE_CONTAINER_PATH"))
	if err != nil {
		http.Error(w, "Invalid contents directory path", http.StatusInternalServerError)
		return
	}

	if !strings.HasPrefix(abs_file_path, abs_contents_dir) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(file_path))
	http.ServeFile(w, r, file_path)
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
	release, err := c.downloadService.GetLatestRelease(name)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.serve(w, r, release, true)
}

func (c *DownloadController) GetLatestCartridge(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	release, err := c.downloadService.GetLatestRelease(name)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.serve(w, r, release, false)
}

func (c *DownloadController) GetSource(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	version := chi.URLParam(r, "version")
	release, err := c.downloadService.GetSpecificRelease(name, version)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.serve(w, r, release, true)
}

func (c *DownloadController) GetCartridge(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	version := chi.URLParam(r, "version")
	release, err := c.downloadService.GetSpecificRelease(name, version)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.serve(w, r, release, false)
}
