package handlers

import (
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/auth"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/role"
)

type handler struct{}

type Handler interface {
	// TODO: handlers
	HealthHandler
	AuthHandler
	RoleHandler
}

// TODO: handlers
type HealthHandler interface {
	GetHealth(rt *rest_goswagger_api.Runtime) (response *models.BasicResponse, err error)
}

type AuthHandler interface {
	PostAuthLogin(rt *rest_goswagger_api.Runtime, params *auth.PostAuthLoginParams) (response *models.PostAuthLoginResponse, err error)
	PostAuthRegistration(rt *rest_goswagger_api.Runtime, params *auth.PostAuthRegistrationParams) (response *models.PostAuthRegistrationResponse, err error)
}

type RoleHandler interface {
	PostRole(rt *rest_goswagger_api.Runtime, params *role.PostRoleParams) (response *models.BasicResponse, err error)
}

func NewHandler() Handler {
	return &handler{}
}
