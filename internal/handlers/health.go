package handlers

import (
	"net/http"

	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
)

func (h *handler) GetHealth(rt *rest_goswagger_api.Runtime) (response *models.BasicResponse, err error) {

	return &models.BasicResponse{
		Code:    http.StatusOK,
		Message: "OK",
	}, nil

}
