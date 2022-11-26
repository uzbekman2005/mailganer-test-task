package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	LogLevel         string
	Environment      string
	PostgresHost     string
	PostgresDatabase string
	PostgresPort     string
	PostgresPassword string
	PostgresUser     string
}

func LoadConfig() *Config {
	c := &Config{}

	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop")) // develop,  staging, production
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABASE", "maildb"))
	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "azizbek"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "Azizbek"))

	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
