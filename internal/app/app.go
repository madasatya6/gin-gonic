// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/madasatya6/gin-gonic/config"
	amqprpc "github.com/madasatya6/gin-gonic/internal/controller/amqp_rpc"
	v1 "github.com/madasatya6/gin-gonic/internal/controller/http/v1"
	"github.com/madasatya6/gin-gonic/internal/usecase"
	"github.com/madasatya6/gin-gonic/internal/usecase/repo"
	"github.com/madasatya6/gin-gonic/internal/usecase/webapi"
	"github.com/madasatya6/gin-gonic/pkg/httpserver"
	"github.com/madasatya6/gin-gonic/pkg/logger"
	"github.com/madasatya6/gin-gonic/pkg/postgres"
	"github.com/madasatya6/gin-gonic/pkg/session"
	"github.com/madasatya6/gin-gonic/pkg/rabbitmq/rmq_rpc/server"
	"github.com/madasatya6/gin-gonic/pkg/funcmap"
	"github.com/madasatya6/gin-gonic/pkg/view"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	translationUseCase := usecase.New(
		repo.New(pg),
		webapi.New(),
	)

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(translationUseCase)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// HTTP Server [route assets]
	gin.ForceConsoleColor()
	handler := gin.New()
	session.New(handler)
	handler.SetFuncMap(funcmap.FuncMap)
	handler.HTMLRender = view.New()
	handler.Static("/assets", "./app/static")

	//routing versi 1
	v1.NewRouter(handler, l, translationUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
