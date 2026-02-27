package http

import (
	"fmt"
	"net/http"
	"os"
	"teletype_softwares/domain"

	"github.com/go-chi/chi/v5"
)

type DocsController struct {
	softwareService domain.SoftwareServiceInterface
	fileService     domain.FileServiceInterface
}

func NewDocsController(software_service domain.SoftwareServiceInterface, file_service domain.FileServiceInterface) *DocsController {
	return &DocsController{
		softwareService: software_service,
		fileService:     file_service,
	}
}

func (c *DocsController) ServeDocs(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	version := chi.URLParam(r, "version")

	software, err := c.softwareService.GetByName(name)
	if err != nil {
		if err == domain.ErrSoftwareNotFound {
			http.Error(w, "Software not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var targetRelease *domain.Release
	for _, release := range software.Releases {
		if release.Version == version && release.DocsFolderPath != "" {
			targetRelease = &release
			break
		}
	}

	if targetRelease == nil {
		http.Error(w, "No documentation found for this release", http.StatusNotFound)
		return
	}

	docs_base_dir := c.fileService.GetPath(targetRelease.DocsFolderPath)

	if _, err := os.Stat(docs_base_dir); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("Documentation for '%s' version '%s' not found.", name, version), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Error accessing documentation for '%s' version '%s': %v", name, version, err), http.StatusInternalServerError)
		return
	}

	fs := http.StripPrefix(fmt.Sprintf("/docs/%s/%s", name, version), http.FileServer(http.Dir(docs_base_dir)))
	fs.ServeHTTP(w, r)
}
