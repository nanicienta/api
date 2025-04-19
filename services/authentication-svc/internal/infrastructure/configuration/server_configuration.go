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

var httpServerConfiguration *HTTPServerConfig

// GetHTTPServerConfig Returns the initialized HTTPServerConfig
func GetHTTPServerConfig() *HTTPServerConfig {
	if httpServerConfiguration == nil {
		initHTTPServerConfig()
	}
	return httpServerConfiguration
}

func initHTTPServerConfig() {
	portStr := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 9003
	}

	httpServerConfiguration = &HTTPServerConfig{
		port: uint(port),
	}
}
