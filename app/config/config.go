package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	LogLevel         string
	Host             string
	Port             string
	Environment      string
	SignInKey        string
	AuthConfigPath   string
	CSVFilePath      string
	RedisHost        string
	RedisPort        string
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
	c.Host = cast.ToString(GetOrReturnDefault("HOST", "localhost"))
	c.Port = cast.ToString(GetOrReturnDefault("PORT", "9090"))

	c.SignInKey = cast.ToString(GetOrReturnDefault("SIGN_IN_KEY", "sdfasfsadfa"))
	c.AuthConfigPath = cast.ToString(GetOrReturnDefault("AUTH_CONFIG_PATH", "/home/azizbek/go/src/github.com/uzbekman2005/mailganer-test-task/app/config/auth.conf"))
	c.CSVFilePath = cast.ToString(GetOrReturnDefault("CSV_FILE_PATH", "/home/azizbek/go/src/github.com/uzbekman2005/mailganer-test-task/app/config/auth.csv"))

	c.RedisHost = cast.ToString(GetOrReturnDefault("REDIS_HOST", "localhost"))
	c.RedisPort = cast.ToString(GetOrReturnDefault("REDIS_PORT", "6379"))

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
