package config

type Config struct {
	AppName string
	Port    string
	DbUrl   string
	DbDebug bool
}

func Load() Config {
	return Config{
		AppName: "user-service",
		Port:    "8080",
		DbUrl:   "postgres://root:root@localhost:54321/user-service?sslmode=disable",
		DbDebug: true,
	}
}
