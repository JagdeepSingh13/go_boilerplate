package handler

import (
	"github.com/JagdeepSingh13/go_boilerplate/internal/server"
	"github.com/JagdeepSingh13/go_boilerplate/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:  NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
	}
}
