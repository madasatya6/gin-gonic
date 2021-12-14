package amqprpc

import (
	"github.com/madasatya6/gin-gonic/internal/usecase"
	"github.com/madasatya6/gin-gonic/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
