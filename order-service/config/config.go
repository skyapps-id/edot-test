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
	JwtSecret                 []byte
	HostProductService        string
	TokenProductService       string
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
		JwtSecret:                 []byte(os.Getenv("JWT_SECRET")),
		HostProductService:        os.Getenv("HOST_PRODUCT_SERVICE"),
		TokenProductService:       os.Getenv("TOKEN_PRODUCT_SERVICE"),
		HostShopWarehouseService:  os.Getenv("HOST_SHOP_WAREHOUSE_SERVICE"),
		TokenShopWarehouseService: os.Getenv("TOKEN_SHOP_WAREHOUSE_SERVICE"),
	}
}
