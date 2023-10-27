package routes

import (
	"github.com/go-openapi/runtime/middleware"
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/app"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/handlers"
)

func setRouteHealth(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI, apiHandler handlers.Handler) {

	api.AppGetHealthHandler = app.GetHealthHandlerFunc(func(ghp app.GetHealthParams) middleware.Responder {
		response, err := apiHandler.GetHealth(rt)
		if err != nil {
			errResponse := rt.GetError(err)
			return app.NewGetHealthDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    errResponse.Code(),
				Message: errResponse.Error(),
			})
		}

		return app.NewGetHealthOK().WithPayload(response)
	})

	// more route here
	// , , ,

}
