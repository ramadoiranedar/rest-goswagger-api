package routes

import (
	"github.com/go-openapi/runtime/middleware"
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/movie"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/handlers"
)

func setRouteMovie(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI, apiHandler handlers.Handler) {

	// GET /movie
	api.MovieGetMovieTrendingHandler = movie.GetMovieTrendingHandlerFunc(func(params movie.GetMovieTrendingParams, p *models.Principal) middleware.Responder {
		response, err := apiHandler.GetMovie(rt, &params)
		if err != nil {
			errResponse := rt.GetError(err)
			return movie.NewGetMovieTrendingDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    models.Code(errResponse.Code()),
				Message: models.Message(errResponse.Error()),
			})
		}

		return movie.NewGetMovieTrendingOK().WithPayload(response)
	})

}
