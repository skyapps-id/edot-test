package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string
	DbUrl   string
	DbDebug bool
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found or failed to load it")
	}

	dbDebug, _ := strconv.ParseBool(os.Getenv("DB_DEBUG"))

	return Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
		DbUrl:   os.Getenv("DB_URL"),
		DbDebug: dbDebug,
	}
}
