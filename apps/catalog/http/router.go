package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	apiSoftwareController            *APISoftwareController
	apiSoftwareHighlightedController *APISoftwareHighlightedController
	softwareUpdaterController        *SoftwareUpdaterController
}

func NewRouter(
	api_software_controller *APISoftwareController,
	api_software_highlighted_controller *APISoftwareHighlightedController,
	software_updater_controller *SoftwareUpdaterController,
) *Router {
	return &Router{
		apiSoftwareController:            api_software_controller,
		apiSoftwareHighlightedController: api_software_highlighted_controller,
		softwareUpdaterController:        software_updater_controller,
	}
}

func (r *Router) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(CORSMiddleware)

	router.Get("/api/software", r.apiSoftwareController.Index)
	router.Get("/api/software/highlighted", r.apiSoftwareHighlightedController.Index)
	router.Get("/update", r.softwareUpdaterController.Update)

	fs_file := http.FileServer(http.Dir("/softwares"))
	router.Handle("/file/*", http.StripPrefix("/file/", fs_file))

	return router
}
