package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	softwareController        *SoftwareController
	softwareUpdaterController *SoftwareUpdaterController
	downloadController        *DownloadController
	playController            *PlayController
	rootController            *RootController
}

func NewRouter(
	softwareController *SoftwareController,
	softwareUpdaterController *SoftwareUpdaterController,
	downloadController *DownloadController,
	playController *PlayController,
	rootController *RootController,
) *Router {
	return &Router{
		softwareController:        softwareController,
		softwareUpdaterController: softwareUpdaterController,
		downloadController:        downloadController,
		playController:            playController,
		rootController:            rootController,
	}
}

func (r *Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", r.rootController.index)
	router.Get("/software", r.softwareController.index)
	router.Get("/software/{name}", r.softwareController.show)

	router.Get("/update", r.softwareUpdaterController.update)
	router.Get("/download/{name}/source", r.downloadController.GetLatestSource)
	router.Get("/download/{name}/cartridge", r.downloadController.GetLatestCartridge)
	router.Get("/download/{name}/{version}/source", r.downloadController.GetSource)
	router.Get("/download/{name}/{version}/cartridge", r.downloadController.GetCartridge)
	router.Get("/play/{name}/{version}", r.playController.Play)
	router.Get("/play/{name}/{version}/content*", r.playController.ServeContent)

	fs := http.FileServer(http.Dir("assets"))
	router.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	return router
}
