package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	apiSoftwareController     *APISoftwareController
	softwareUpdaterController *SoftwareUpdaterController
}

func NewRouter(
	api_software_controller *APISoftwareController,
	software_updater_controller *SoftwareUpdaterController,
) *Router {
	return &Router{
		apiSoftwareController:     api_software_controller,
		softwareUpdaterController: software_updater_controller,
	}
}

func (r *Router) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(CORSMiddleware)

	router.Get("/api/software", r.apiSoftwareController.Index)
	router.Get("/update", r.softwareUpdaterController.Update)

	fs_file := http.FileServer(http.Dir("/softwares"))
	router.Handle("/file/*", http.StripPrefix("/file/", fs_file))

	return router
}
