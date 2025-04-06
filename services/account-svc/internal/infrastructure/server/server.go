// Package server defines the elements used for the server. (In this case HTTP)
package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nanicienta/api/account-svc/internal/infrastructure/configuration"
	"log"
	"net/http"
	"time"
)

const defaultHost = "0.0.0.0"

// HTTPServer Interface for the HTTP server
type HTTPServer interface {
	Start()
	Stop()
}

type httpServer struct {
	Port   uint
	server *http.Server
}

// NewHTTPServer Creates a new HTTPServer
func NewHTTPServer(
	router *gin.Engine,
	config *configuration.HTTPServerConfig,
) HTTPServer {
	return &httpServer{
		Port: config.GetPort(),
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", defaultHost, config.GetPort()),
			Handler: router,
		},
	}
}

func (httpServer httpServer) Start() {
	go func() {
		if err := httpServer.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(
				"failed to stater HTTPServer listen port %d, err=%s",
				httpServer.Port, err.Error(),
			)
		}
	}()

	log.Printf("Start Service with port %d", httpServer.Port)
}

func (httpServer httpServer) Stop() {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(3)*time.Second,
	)
	defer cancel()

	if err := httpServer.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown err=%s", err.Error())
	}
}
