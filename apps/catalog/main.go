package main

import (
	"teletype_softwares/domain"
	"teletype_softwares/http"
)

func main() {
	domain := domain.NewDomain()
	http.StartHttpServer(domain)
}
