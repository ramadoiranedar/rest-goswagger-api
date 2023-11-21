package utils

import "net/http"

type customTransport struct {
	originalTransport http.RoundTripper
}

func NewCustomTransport() *customTransport {
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
