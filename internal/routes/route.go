package routes

import (
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/handlers"
)

func Route(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI, apiHandler handlers.Handler) {
	rt.Logger().Info("Initialize Route")

	// set all route here
	setRouteHealth(rt, api, apiHandler)
	setRouteAuth(rt, api, apiHandler)

}
