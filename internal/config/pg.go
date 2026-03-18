package config

import "fmt"

type PGConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	SSLMode  string
}

func LoadPGConfig() PGConfig {
	return PGConfig{
		Host:     mustEnv("PG_HOST"),
		Port:     getEnv("PG_PORT", "55555"),
		DBName:   mustEnv("PG_DATABASE_NAME"),
		User:     mustEnv("PG_USER"),
		Password: mustEnv("PG_PASSWORD"),
		SSLMode:  getEnv("PG_SSL_MODE", "disable"),
	}
}

func (p PGConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		p.Host, p.Port, p.DBName, p.User, p.Password, p.SSLMode,
	)
}
