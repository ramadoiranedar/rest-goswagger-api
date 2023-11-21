package utils

import (
	"context"
	"fmt"
	"time"

	themoviedb_service "github.com/ramadoiranedar/rest-goswagger-api/pkg/client/themoviedb_client"

	"github.com/go-openapi/runtime"
	http_transport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
)

func GetTheMovieDbClientSDK(url string) *themoviedb_service.Themoviedb {
	transport := http_transport.New(url, "", []string{"https"})
	transport.Transport = NewCustomTransport()

	transport.Debug = true // TODO: FOR ENV DEVELOPMENT

	return themoviedb_service.New(transport, strfmt.Default)
}

func GetTheMovieDbUrl(rt *rest_goswagger_api.Runtime, url string) string {
	return url
}

func GetContextTimeout(parent context.Context, rt *rest_goswagger_api.Runtime) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, time.Second*time.Duration(rt.Conf.GetInt64("themoviedb.timeout")))
}

func GetTheMovieDbAuth(token string) runtime.ClientAuthInfoWriter {
	return http_transport.APIKeyAuth("Authorization", "header", fmt.Sprintf("Bearer %s", token))
}
