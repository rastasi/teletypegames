package http

import (
	"teletype_softwares/domain"
	"teletype_softwares/lib/http_utils"
)

func StartHttpServer(domain_instance domain.Domain) {
	router := NewRouter(
		NewSoftwareController(domain_instance.SoftwareService),
		NewSoftwareUpdaterController(domain_instance.SoftwareUpdaterService, domain_instance.SoftwareService),
		NewDownloadController(domain_instance.DownloadService, domain_instance.SoftwareService),
		NewPlayController(domain_instance.SoftwareService, domain_instance.FileService),
		NewRootController(domain_instance.SoftwareService),
		NewDocsController(domain_instance.SoftwareService, domain_instance.FileService),
	).Init()

	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router: router,
	})
}
