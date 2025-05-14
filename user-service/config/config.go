package config

type Config struct {
	DbUrl   string
	DbDebug bool
}

func Load() Config {
	return Config{
		DbUrl:   "postgres://root:root@localhost:54321/user-service?sslmode=disable",
		DbDebug: true,
	}
}
