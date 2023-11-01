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

}
