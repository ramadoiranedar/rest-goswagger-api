package routes

import (
	"github.com/go-openapi/runtime/middleware"
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/auth"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/handlers"
)

func setRouteAuth(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI, apiHandler handlers.Handler) {

	api.AuthPostAuthLoginHandler = auth.PostAuthLoginHandlerFunc(func(params auth.PostAuthLoginParams) middleware.Responder {
		response, err := apiHandler.PostAuthLogin(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return auth.NewPostAuthLoginDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return auth.NewPostAuthLoginOK().WithPayload(response)
	})

	api.AuthPostAuthRegistrationHandler = auth.PostAuthRegistrationHandlerFunc(func(params auth.PostAuthRegistrationParams) middleware.Responder {
		response, err := apiHandler.PostAuthRegistration(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return auth.NewPostAuthRegistrationDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return auth.NewPostAuthRegistrationCreated().WithPayload(response)
	})

}
