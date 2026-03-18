package config

import (
	"log"
	"os"
)

type Config struct {
	Env  string
	GRPC GRPCConfig
	PG   PGConfig
}

func LoadConfig() Config {
	return Config{
		Env:  getEnv("ENV", "local"),
		GRPC: LoadGRPCConfig(),
		PG:   LoadPGConfig(),
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
