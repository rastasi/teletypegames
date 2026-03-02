package http

import (
	"teletype_softwares/domain"
	"teletype_softwares/lib/http_utils"
)

func StartHttpServer(domain_instance domain.Domain) {
	router := NewRouter(
		NewAPISoftwareController(domain_instance.SoftwareService),
		NewSoftwareUpdaterController(domain_instance.SoftwareUpdaterService, domain_instance.SoftwareService),
	).Init()

	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router: router,
	})
}
