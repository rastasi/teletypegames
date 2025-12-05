package http

import (
	"teletype_softwares/domain"
	"teletype_softwares/lib/http_utils"
)

func StartHttpServer(domain domain.Domain) {
	router := Router{
		SoftwareController: &SoftwareController{softwareService: domain.SoftwareService},
		SoftwareUpdaterController: &SoftwareUpdaterController{
			softwareUpdaterService: domain.SoftwareUpdaterService,
		},
		DownloadController: &DownloadController{downloadService: domain.DownloadService},
		PlayController:     &PlayController{},
	}.Init()

	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router: router,
	})
}
