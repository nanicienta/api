// Package configuration provides the configuration for the different components
package configuration

import (
	"os"
	"strconv"
)

// HTTPServerConfig holds the configuration for the HTTP server
type HTTPServerConfig struct {
	port uint
}

// GetPort returns the port on which the HTTP server will listen
func (httpServerConfig *HTTPServerConfig) GetPort() uint {
	return httpServerConfig.port
}

var initHTTPServerConfig *HTTPServerConfig

// GetHTTPServerConfig Returns the initialized HTTPServerConfig
func GetHTTPServerConfig() *HTTPServerConfig {
	if initHTTPServerConfig == nil {
		initConfig()
	}
	return initHTTPServerConfig
}

func initConfig() {
	portStr := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 9000
	}

	initHTTPServerConfig = &HTTPServerConfig{
		port: uint(port),
	}
}
