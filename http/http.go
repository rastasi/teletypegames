package http

import (
	"teletype_softwares/domain"
	"teletype_softwares/lib/http_utils"
)

func StartHttpServer(domain_instance domain.Domain) {
	router := NewRouter(
		NewSoftwareController(domain_instance.SoftwareService),
		NewSoftwareUpdaterController(domain_instance.SoftwareUpdaterService),
		NewDownloadController(domain_instance.DownloadService),
		NewPlayController(domain_instance.SoftwareService, domain_instance.FileService),
		NewRootController(),
	).Init()

	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router: router,
	})
}
