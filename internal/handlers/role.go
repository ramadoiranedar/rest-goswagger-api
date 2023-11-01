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

	response = &models.BasicResponse{
		Code:    http.StatusCreated,
		Message: constants.MSG_CREATED,
	}

	return
}

func (h *handler) PutRoleID(rt *rest_goswagger_api.Runtime, params *role.PutRoleIDParams) (response *models.BasicResponse, err error) {
	role := models.Role{
		ModelIdentifier: models.ModelIdentifier{
			ID: &params.ID,
		},
		DataRole: models.DataRole{
			Name: params.Name,
			Slug: params.Slug,
		},
	}

	if err = rt.DbMysql.Updates(&role).Error; err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	response = &models.BasicResponse{
		Code:    http.StatusOK,
		Message: constants.MSG_OK,
	}

	return
}

func (h *handler) DeleteRoleID(rt *rest_goswagger_api.Runtime, params *role.DeleteRoleIDParams) (response *models.BasicResponse, err error) {
	role := models.Role{
		ModelIdentifier: models.ModelIdentifier{
			ID: &params.ID,
		},
	}

	db := rt.DbMysql.Delete(&role)
	if err = db.Error; err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}
	if db.RowsAffected < 1 {
		err = rt.SetError(http.StatusUnprocessableEntity, constants.MSG_ROLE_NOT_EXISTS)
		return
	}

	response = &models.BasicResponse{
		Code:    http.StatusOK,
		Message: constants.MSG_OK,
	}

	return
}

func (h *handler) GetRoleID(rt *rest_goswagger_api.Runtime, params *role.GetRoleIDParams) (response *models.GetRoleIDResponse, err error) {
	role := models.Role{}

	if err = rt.DbMysql.Where("id = ?", params.ID).First(&role).Error; err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	response = &models.GetRoleIDResponse{
		Code:    http.StatusOK,
		Message: constants.MSG_OK,
		Data:    &role,
	}

	return
}

func (h *handler) GetRole(rt *rest_goswagger_api.Runtime, params *role.GetRoleParams) (response *models.GetRoleResponse, err error) {
	var (
		role = []*models.Role{}
		db   = rt.DbMysql.Model(&role)
	)

	if params.Sort != nil && params.SortBy != nil {
		orderBy := fmt.Sprintf("%s %s", *params.SortBy, *params.Sort)
		db = db.Order(orderBy)
	}

	if err = db.Find(&role).Error; err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	response = &models.GetRoleResponse{
		Code:    http.StatusOK,
		Message: constants.MSG_OK,
		Data:    role,
	}

	return
}
