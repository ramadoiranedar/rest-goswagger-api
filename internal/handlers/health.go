package handlers

import rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"

func (h *handler) GetHealth(rt *rest_goswagger_api.Runtime) string {
	return "server is running"
}
