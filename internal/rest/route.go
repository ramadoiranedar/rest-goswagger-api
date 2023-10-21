package rest

import (
	"net/http"

	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/app"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/handlers"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI, apiHandler handlers.Handler) {
	rt.Logger().Info("Initialize Route")

	api.AppGetHealthHandler = app.GetHealthHandlerFunc(func(ghp app.GetHealthParams, p *models.Principal) middleware.Responder {
		result := apiHandler.GetHealth(rt)
		return app.NewGetHealthOK().WithPayload(&models.DefaultResponse{
			Code:    http.StatusOK,
			Message: result,
		})
	})
}
