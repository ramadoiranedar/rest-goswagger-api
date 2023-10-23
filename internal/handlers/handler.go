package handlers

import (
	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
)

type handler struct{}

type Handler interface {
	// TODO: handlers
	HealthHandler
}

// TODO: handlers
type HealthHandler interface {
	GetHealth(rt *rest_goswagger_api.Runtime) (response *models.BasicResponse, err error)
}

func NewHandler() Handler {
	return &handler{}
}
