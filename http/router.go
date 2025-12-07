package http

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	softwareController        *SoftwareController
	softwareUpdaterController *SoftwareUpdaterController
	downloadController        *DownloadController
	playController            *PlayController
}

func NewRouter(
	softwareController *SoftwareController,
	softwareUpdaterController *SoftwareUpdaterController,
	downloadController *DownloadController,
	playController *PlayController,
) *Router {
	return &Router{
		softwareController:        softwareController,
		softwareUpdaterController: softwareUpdaterController,
		downloadController:        downloadController,
		playController:            playController,
	}
}

func (r *Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", r.softwareController.index)
	router.Get("/update", r.softwareUpdaterController.update)
	router.Get("/releases/{name}", r.softwareController.releases)
	router.Get("/download/{name}/source", r.downloadController.GetLatestSource)
	router.Get("/download/{name}/cartridge", r.downloadController.GetLatestCartridge)
	router.Get("/download/{name}/{version}/source", r.downloadController.GetSource)
	router.Get("/download/{name}/{version}/cartridge", r.downloadController.GetCartridge)
	router.Get("/play/{name}", r.playController.Play)
	router.Get("/play/{name}/content*", r.playController.ServeContent)

	return router
}
