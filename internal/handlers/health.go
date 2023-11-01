package handlers

import (
	"net/http"

	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/constants"
)

func (h *handler) GetHealth(rt *rest_goswagger_api.Runtime) (response *models.BasicResponse, err error) {

	return &models.BasicResponse{
		Code:    http.StatusOK,
		Message: constants.MSG_OK,
	}, nil

}
