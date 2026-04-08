package config

import (
	"errors"
	"net"
	"os"
)

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	Host string
	Port string
}

func (h *httpConfig) Address() string {
	return net.JoinHostPort(h.Host, h.Port)
}

func NewHTTPConfig() (HTTPConfig,error) {
	host := os.Getenv("HTTP_HOST")
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}
	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &httpConfig{
		Host: host,
		Port: port,
	},nil

}