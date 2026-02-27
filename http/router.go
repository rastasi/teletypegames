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
	docsController            *DocsController
}

func NewRouter(
	software_controller *SoftwareController,
	software_updater_controller *SoftwareUpdaterController,
	download_controller *DownloadController,
	play_controller *PlayController,
	root_controller *RootController,
	docs_controller *DocsController,
) *Router {
	return &Router{
		softwareController:        software_controller,
		softwareUpdaterController: software_updater_controller,
		downloadController:        download_controller,
		playController:            play_controller,
		rootController:            root_controller,
		docsController:            docs_controller,
	}
}

func (r *Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", r.rootController.Index)
	router.Get("/software", r.softwareController.Index)
	router.Get("/software/{name}", r.softwareController.Show)

	router.Get("/update", r.softwareUpdaterController.Update)
	router.Get("/download/{name}/source", r.downloadController.GetLatestSource)
	router.Get("/download/{name}/cartridge", r.downloadController.GetLatestCartridge)
	router.Get("/download/{name}/{version}/source", r.downloadController.GetSource)
	router.Get("/download/{name}/{version}/cartridge", r.downloadController.GetCartridge)
	router.Get("/play/{name}/{version}", r.playController.Play)
	router.Get("/play/{name}/{version}/content*", r.playController.ServeContent)
	router.Get("/docs/{name}/{version}", r.docsController.ServeDocs)
	router.Get("/docs/{name}/{version}/*", r.docsController.ServeDocs)

	fs := http.FileServer(http.Dir("assets"))
	router.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	return router
}
