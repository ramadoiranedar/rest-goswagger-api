package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	themoviedb_service "github.com/ramadoiranedar/rest-goswagger-api/pkg/client/themoviedb_client"

	"github.com/go-openapi/runtime"
	http_transport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
)

type customTransport struct {
	originalTransport http.RoundTripper
}

func newCustomTransport() *customTransport {
	return &customTransport{
		originalTransport: http.DefaultTransport,
	}
}

func (ct *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := ct.originalTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetInternalTheMovieDbClientSDK(url string) *themoviedb_service.Themoviedb {
	transport := http_transport.New(url, "", []string{"https"})
	transport.Transport = newCustomTransport()

	transport.Debug = true // TODO: FOR DEVELOPMENT

	return themoviedb_service.New(transport, strfmt.Default)
}

func GetInternalTheMovieDbUrl(rt *rest_goswagger_api.Runtime, url string) string {
	return url
}

func GetContextTimeout(parent context.Context, rt *rest_goswagger_api.Runtime) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, time.Second*time.Duration(rt.Conf.GetInt64("themoviedb.timeout")))
}

func GetInternalTheMovieDbAuth(token string) runtime.ClientAuthInfoWriter {
	return http_transport.APIKeyAuth("Authorization", "header", fmt.Sprintf("Bearer %s", token))
}
