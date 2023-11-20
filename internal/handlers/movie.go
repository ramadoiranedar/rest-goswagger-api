package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations/movie"
	utils "github.com/ramadoiranedar/rest-goswagger-api/internal/utils/external_service"
	"github.com/ramadoiranedar/rest-goswagger-api/pkg/client/themoviedb_client/trending"
)

func (h *handler) GetMovie(rt *rest_goswagger_api.Runtime, params *movie.GetMovieTrendingParams) (response *models.GetListTrendingMoviesResponse, err error) {

	// request to service otp whatsapp
	theMovieDbClient := utils.GetInternalTheMovieDbClientSDK(utils.GetInternalTheMovieDbUrl(rt, rt.Conf.GetString("themoviedb.base_url")))
	ctx, cancel := utils.GetContextTimeout(context.Background(), rt)
	defer cancel()

	responseTheMovieDbService, err := theMovieDbClient.Trending.Get3TrendingMovieTimeWindow(&trending.Get3TrendingMovieTimeWindowParams{
		TimeWindow: "week",
		Context:    ctx,
	}, utils.GetInternalTheMovieDbAuth(rt.Conf.GetString("themoviedb.authorization")))
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return nil, err
	}

	byteData, err := json.Marshal(responseTheMovieDbService.Payload)
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return nil, err
	}

	var result *models.GetListTrendingMoviesResponse
	err = json.Unmarshal(byteData, &result)
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return nil, err
	}

	response = result

	return
}
