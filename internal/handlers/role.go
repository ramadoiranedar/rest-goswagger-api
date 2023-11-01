package handlers

import (
	"fmt"
	"net/http"

	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/role"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/constants"
)

func (h *handler) PostRole(rt *rest_goswagger_api.Runtime, params *role.PostRoleParams) (response *models.BasicResponse, err error) {
	role := models.Role{
		DataRole: models.DataRole{
			Name: params.Name,
			Slug: params.Slug,
		},
	}

	if err = rt.DbMysql.Create(&role).Error; err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	message := models.Message(fmt.Sprintf(constants.MSG_ROLE_CREATED, role.Slug))
	response = &models.BasicResponse{
		Code:    http.StatusCreated,
		Message: message,
	}

	return
}
