package http

import (
	"teletype_softwares/domain"
	"teletype_softwares/lib/http_utils"
)

func StartHttpServer(domainInstance domain.Domain) {
	router := NewRouter(
		NewSoftwareController(domainInstance.SoftwareService),
		NewSoftwareUpdaterController(domainInstance.SoftwareUpdaterService),
		NewDownloadController(domainInstance.DownloadService),
		NewPlayController(domainInstance.SoftwareService),
		NewRootController(),
	).Init()

	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router: router,
	})
}
