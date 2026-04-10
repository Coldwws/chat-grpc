package config

type AuthConfig struct {
	Addr string
}

func LoadAuthConfig() AuthConfig {
	return AuthConfig{
		Addr: getEnv("AUTH_GRPC_ADDR", "localhost:50051"),
	}
}