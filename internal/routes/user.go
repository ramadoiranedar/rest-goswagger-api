package routes

import (
	"github.com/go-openapi/runtime/middleware"
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/user"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/handlers"
)

func setRouteUser(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI, apiHandler handlers.Handler) {

	// POST /user
	api.UserPostUserHandler = user.PostUserHandlerFunc(func(params user.PostUserParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.PostUser(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return user.NewPostUserDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return user.NewPostUserCreated().WithPayload(response)
	})

	// PUT /user/{id}
	api.UserPutUserIDHandler = user.PutUserIDHandlerFunc(func(params user.PutUserIDParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.PutUserID(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return user.NewPutUserIDDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return user.NewPutUserIDOK().WithPayload(response)
	})

	// DELETE /user/{id}
	api.UserDeleteUserIDHandler = user.DeleteUserIDHandlerFunc(func(params user.DeleteUserIDParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.DeleteUserID(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return user.NewDeleteUserIDDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return user.NewDeleteUserIDOK().WithPayload(response)
	})

	// GET /user/{id}
	api.UserGetUserIDHandler = user.GetUserIDHandlerFunc(func(params user.GetUserIDParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.GetUserID(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return user.NewGetUserIDDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return user.NewGetUserIDOK().WithPayload(response)
	})

	// // GET /user
	api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.GetUser(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return user.NewGetUserDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return user.NewGetUserOK().WithPayload(response)
	})

}
