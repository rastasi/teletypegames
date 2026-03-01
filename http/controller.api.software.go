package http

import (
	"encoding/json"
	"net/http"

	"teletype_softwares/domain"
)

type APISoftwareController struct {
	softwareService domain.SoftwareServiceInterface
}

func NewAPISoftwareController(software_service domain.SoftwareServiceInterface) *APISoftwareController {
	return &APISoftwareController{softwareService: software_service}
}

func (c *APISoftwareController) Index(w http.ResponseWriter, r *http.Request) {
	softwares, _ := c.softwareService.DetailedList()
	json.NewEncoder(w).Encode(softwares)
}
