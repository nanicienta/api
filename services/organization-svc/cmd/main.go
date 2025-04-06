// Application service entry
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nanicienta/api/organization-svc/internal/infrastructure/adapters/logging"
	"github.com/nanicienta/api/organization-svc/internal/infrastructure/configuration"
	"github.com/nanicienta/api/organization-svc/internal/infrastructure/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
	//Initialize the logger
	log := logging.InitZapLogger()
	log.Info("Starting account service")

	// Initialize the Server configuration
	instance := gin.New()
	instance.Use(gin.Recovery())
	httpServer := server.NewHTTPServer(
		instance,
		configuration.GetHTTPServerConfig(),
	)
	httpServer.Start()
	defer httpServer.Stop()
	log.Info("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Info("graceful shutdown...")
}
