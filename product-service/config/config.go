package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName                   string
	Port                      string
	HostOTLP                  string
	DbUrl                     string
	DbDebug                   bool
	TokenInternal             string
	JwtSecret                 []byte
	HostShopWarehouseService  string
	TokenShopWarehouseService string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found or failed to load it")
	}

	dbDebug, _ := strconv.ParseBool(os.Getenv("DB_DEBUG"))

	return Config{
		AppName:                   os.Getenv("APP_NAME"),
		Port:                      os.Getenv("PORT"),
		HostOTLP:                  os.Getenv("HOST_OTLP"),
		DbUrl:                     os.Getenv("DB_URL"),
		DbDebug:                   dbDebug,
		TokenInternal:             os.Getenv("TOKEN_INTERNAL"),
		JwtSecret:                 []byte(os.Getenv("JWT_SECRET")),
		HostShopWarehouseService:  os.Getenv("HOST_SHOP_WAREHOUSE_SERVICE"),
		TokenShopWarehouseService: os.Getenv("TOKEN_SHOP_WAREHOUSE_SERVICE"),
	}
}
