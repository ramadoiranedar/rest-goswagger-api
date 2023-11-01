package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/auth"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/constants"
	utils_hash "github.com/ramadoiranedar/rest-goswagger-api/internal/utils/hash"
	utils_jwt "github.com/ramadoiranedar/rest-goswagger-api/internal/utils/jwt"
	"gorm.io/gorm"
)

func (h *handler) PostAuthLogin(rt *rest_goswagger_api.Runtime, params *auth.PostAuthLoginParams) (response *models.PostAuthLoginResponse, err error) {
	user := models.User{}

	err = rt.DbMysql.Model(user).
		Preload("Role").
		Where("email = ?", params.Email).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = rt.SetError(http.StatusNotFound, constants.MSG_USER_NOT_FOUND)
		}
		return
	}

	passwordHash := utils_hash.HashSha256(*params.Password)
	if *user.Password != passwordHash {
		err = rt.SetError(http.StatusUnprocessableEntity, constants.MSG_INCORRECT_EMAIL_OR_PASSWORD)
		return
	}

	maker, err := utils_jwt.NewJWTMaker(rt.Conf.GetString("jwt.secret"))
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	expiryToken := time.Duration(rt.Conf.GetInt("jwt.expired")) * time.Second
	paramPayload := &utils_jwt.PayloadJWT{
		UserID: *user.ID,
		Email:  strfmt.Email(*user.Email),
	}
	if user.Role != nil {
		paramPayload.RoleSlug = user.Role.Slug
	}

	token, payload, err := maker.CreateToken(paramPayload, expiryToken)
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	data := &models.PostAuthLoginDataResponse{
		Token: token,
		PayloadJwt: &models.PayloadJwt{
			UserID:    payload.UserID,
			Email:     payload.Email,
			RoleSlug:  payload.RoleSlug,
			ExpiredAt: strfmt.DateTime(payload.ExpiredAt),
			IssuedAt:  strfmt.DateTime(payload.IssuedAt),
		},
	}

	response = &models.PostAuthLoginResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	}

	return

}

func (h *handler) PostAuthRegistration(rt *rest_goswagger_api.Runtime, params *auth.PostAuthRegistrationParams) (response *models.PostAuthRegistrationResponse, err error) {

	passwordHash := utils_hash.HashSha256(params.Password)
	user := &models.User{
		DataUser: models.DataUser{
			Email:    (*string)(&params.Email),
			Fullname: &params.Fullname,
			Password: &passwordHash,
			RoleID:   constants.ROLE_ID_ADMIN,
		},
	}

	if err = rt.DbMysql.Model(&user).Preload("Role").Create(user).Error; err != nil {
		errCode := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			errCode = http.StatusBadRequest
			err = fmt.Errorf(constants.MSG_EMAIL_IS_REGISTERED)
		}
		err = rt.SetError(errCode, err.Error())
		return
	}

	msg := fmt.Sprintf(constants.MSG_SUCCESS_REGISTRATION, models.Message(*user.Email))
	response = &models.PostAuthRegistrationResponse{
		Code:    http.StatusCreated,
		Message: models.Message(msg),
	}

	return

}
