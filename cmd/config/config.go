package config

import (
	"os"
	"strconv"
)

type DataBaseConfig struct {
	Address string
	Name    string
}

type AppConfig struct {
	Address   string
	Debug     bool
	SecretKey string
}

type Config struct {
	DataBase DataBaseConfig
	App      AppConfig
}

func New() *Config {
	return &Config{
		DataBase: DataBaseConfig{
			Address: getEnv("database_url", ""),
			Name:    getEnv("database_name", ""),
		},
		App: AppConfig{
			Address:   getEnv("address", ""),
			Debug:     getEnvAsBool("debug", false),
			SecretKey: getEnv("secret_key", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	valStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valStr); err != nil {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := getEnv(key, "")
	if value, err := strconv.Atoi(valStr); err != nil {
		return value
	}
	return defaultVal
}
