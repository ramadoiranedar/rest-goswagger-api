package handlers

import (
	"fmt"
	"net/http"
	"strings"

	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/user"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/constants"
	utils_hash "github.com/ramadoiranedar/rest-goswagger-api/internal/utils/hash"
	utils_query "github.com/ramadoiranedar/rest-goswagger-api/internal/utils/query"
)

func (h *handler) PostUser(rt *rest_goswagger_api.Runtime, params *user.PostUserParams) (response *models.BasicResponse, err error) {
	passwordHash := utils_hash.HashSha256(params.Password)
	user := models.User{
		DataUser: models.DataUser{
			Email:    (*string)(&params.Email),
			Password: &passwordHash,
			Fullname: &params.Fullname,
			RoleID:   constants.MAP_ROLE_ID[uint64(params.AuthType)],
		},
	}

	if err = rt.DbMysql.Create(&user).Error; err != nil {
		errCode := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			errCode = http.StatusBadRequest
			err = fmt.Errorf(constants.MSG_EMAIL_IS_REGISTERED)
		}
		err = rt.SetError(errCode, err.Error())
		return
	}

	response = &models.BasicResponse{
		Code:    http.StatusCreated,
		Message: constants.MSG_CREATED,
	}

	return
}

func (h *handler) PutUserID(rt *rest_goswagger_api.Runtime, params *user.PutUserIDParams) (response *models.BasicResponse, err error) {
	passwordHash := utils_hash.HashSha256(params.Password)
	user := models.User{
		ModelIdentifier: models.ModelIdentifier{
			ID: &params.ID,
		},
		DataUser: models.DataUser{
			Email:    (*string)(&params.Email),
			Password: &passwordHash,
			Fullname: &params.Fullname,
			RoleID:   constants.MAP_ROLE_ID[uint64(params.AuthType)],
		},
	}

	if err = rt.DbMysql.Updates(&user).Error; err != nil {
		errCode := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			errCode = http.StatusBadRequest
			err = fmt.Errorf(constants.MSG_EMAIL_IS_REGISTERED)
		}
		err = rt.SetError(errCode, err.Error())
		return
	}

	response = &models.BasicResponse{
		Code:    http.StatusOK,
		Message: constants.MSG_OK,
	}

	return
}

func (h *handler) DeleteUserID(rt *rest_goswagger_api.Runtime, params *user.DeleteUserIDParams) (response *models.BasicResponse, err error) {
	user := models.User{
		ModelIdentifier: models.ModelIdentifier{
			ID: &params.ID,
		},
	}

	db := rt.DbMysql.Delete(&user)
	if err = db.Error; err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}
	if db.RowsAffected < 1 {
		err = rt.SetError(http.StatusUnprocessableEntity, fmt.Sprintf(constants.MSG_NOT_EXISTS, "user"))
		return
	}

	response = &models.BasicResponse{
		Code:    http.StatusOK,
		Message: constants.MSG_OK,
	}

	return
}

func (h *handler) GetUserID(rt *rest_goswagger_api.Runtime, params *user.GetUserIDParams) (response *models.GetUserIDResponse, err error) {
	var (
		user = models.User{}
		raw  = `SELECT id, created_at, deleted_at, updated_at, email, fullname FROM users WHERE id = ?`
	)

	if params.WithPassword {
		raw = `SELECT * FROM users WHERE id = ?`
	}

	if err = rt.DbMysql.Raw(raw, params.ID).First(&user).Error; err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	response = &models.GetUserIDResponse{
		Code:    http.StatusOK,
		Message: constants.MSG_OK,
		Data:    &user,
	}

	return
}

func (h *handler) GetUser(rt *rest_goswagger_api.Runtime, params *user.GetUserParams) (response *models.GetUserResponse, err error) {
	var (
		user       = []*models.UserExcludePassword{}
		db         = *rt.DbMysql.Model(&user).Table("users")
		pagination *models.Pagination
	)

	utils_query.ApplySorting(&db, params.Sort, params.SortBy)
	err = utils_query.ApplySearch(rt, &db, params.SearchBy, params.SearchKeyword)
	if err != nil {
		return
	}
	pagination, err = utils_query.ApplyPagination(&db, *params.CurrentPage, *params.PerPage, &user)
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	if err = db.Find(&user).Error; err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	response = &models.GetUserResponse{
		Code:    http.StatusOK,
		Message: constants.MSG_OK,
		Data:    user,
		Metadata: &models.GetUserResponseMetadata{
			Pagination: pagination,
		},
	}

	return
}
