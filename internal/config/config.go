// Package config provides env configuration settings for the application.
package config

import (
	"log"
	"os"
)

type Config struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	Addr string
}

func NewConfig() *Config {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		log.Println("WARNING: LISTEN_ADDR not set")
	}
	return &Config{
		Addr: addr,
	}
}
