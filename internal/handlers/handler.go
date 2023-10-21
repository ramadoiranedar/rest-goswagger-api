package handlers

import rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"

type Handler interface {
	// TODO: handlers
	HealthHandler
}

// TODO: handlers
type HealthHandler interface {
	GetHealth(rt *rest_goswagger_api.Runtime) string
}

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}
