package config

import (
	"log"
	"os"
)

type Config struct {
	Env  string
	GRPC GRPCConfig
	PG   PGConfig
	Http HTTPConfig
	Auth AuthConfig
}

func LoadConfig() Config {
	httpConf,err := NewHTTPConfig()
	if err != nil {
		log.Fatalf("failed to load http config: %s", err)
	}

	return Config{
		Env:  getEnv("ENV", "local"),
		GRPC: LoadGRPCConfig(),
		PG:   LoadPGConfig(),
		Http: httpConf,
		Auth: LoadAuthConfig(),
	}
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s env variable not set!", key)
	}
	return v

}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
