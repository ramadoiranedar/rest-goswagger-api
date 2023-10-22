package rest

import (
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/handlers"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/routes"
)

func Route(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI, apiHandler handlers.Handler) {
	rt.Logger().Info("Initialize Route")

	// set all route here
	routes.SetRouteHealth(rt, api, apiHandler)

}
