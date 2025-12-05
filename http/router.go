package http

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	SoftwareController        *SoftwareController
	SoftwareUpdaterController *SoftwareUpdaterController
	DownloadController        *DownloadController
	PlayController            *PlayController
}

func (r Router) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", r.SoftwareController.index)
	router.Get("/update", r.SoftwareUpdaterController.update)
	router.Get("/releases/{name}", r.SoftwareController.releases)
	router.Get("/download/{name}/source", r.DownloadController.DownloadSource)
	router.Get("/download/{name}/cartridge", r.DownloadController.DownloadCartridge)
	router.Get("/download/{name}/{version}/source", r.DownloadController.DownloadSourceByVersion)
	router.Get("/download/{name}/{version}/cartridge", r.DownloadController.DownloadCartridgeByVersion)
	router.Get("/play/{name}", r.PlayController.PlayGame)
	router.Get("/play/{name}/content*", r.PlayController.ServeGameContent)

	return router
}
