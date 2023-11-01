package routes

import (
	"github.com/go-openapi/runtime/middleware"
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/role"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/handlers"
)

func setRouteRole(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI, apiHandler handlers.Handler) {

	// POST /role
	api.RolePostRoleHandler = role.PostRoleHandlerFunc(func(params role.PostRoleParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.PostRole(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return role.NewPostRoleDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return role.NewPostRoleCreated().WithPayload(response)
	})

	// PUT /role/{id}
	api.RolePutRoleIDHandler = role.PutRoleIDHandlerFunc(func(params role.PutRoleIDParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.PutRoleID(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return role.NewPutRoleIDDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return role.NewPutRoleIDOK().WithPayload(response)
	})

	// DELETE /role/{id}
	api.RoleDeleteRoleIDHandler = role.DeleteRoleIDHandlerFunc(func(params role.DeleteRoleIDParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.DeleteRoleID(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return role.NewDeleteRoleIDDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return role.NewDeleteRoleIDOK().WithPayload(response)
	})

	// GET /role/{id}
	api.RoleGetRoleIDHandler = role.GetRoleIDHandlerFunc(func(params role.GetRoleIDParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.GetRoleID(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return role.NewGetRoleIDDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return role.NewGetRoleIDOK().WithPayload(response)
	})

	// GET /role
	api.RoleGetRoleHandler = role.GetRoleHandlerFunc(func(params role.GetRoleParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.GetRole(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return role.NewGetRoleDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return role.NewGetRoleOK().WithPayload(response)
	})

}
