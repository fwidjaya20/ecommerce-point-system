package config

import "os"

// noinspection ALL
const (
	ENV 			= "ENV"
	ENV_DEVELOPMENT = "development"

	SERVICE_NAME string = "SERVICE_NAME"

	HTTP_ADDR = "HTTP_ADDR"
	NATS_ADDR = "NATS_ADDR"

	DB_DRIVER = "DB_DRIVER"
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"
	DB_NAME = "DB_NAME"

	MIGRATION_PATH = "MIGRATION_PATH"
)

var defaultConfig = map[string]string{
	// Common Configuration
	ENV:          ENV_DEVELOPMENT,
	SERVICE_NAME: "ecommerce-point-system",

	// Database Configuration
	DB_DRIVER: "postgres",
	DB_HOST:   "localhost",
	DB_PORT:   "5432",
	DB_NAME:   "ecommerce_point_system",
	DB_USER:   "postgres",
	DB_PASS:   "password",

	// Migration and Seeder
	MIGRATION_PATH: "internal/databases/migrations",

	// Transport
	HTTP_ADDR: ":8001",
	NATS_ADDR: ":4222",
}

func GetEnv(key string) string {
	r := os.Getenv(key)

	if r == "" {
		if _, ok := defaultConfig[key]; !ok {
			return ""
		}
		r = defaultConfig[key]
	}

	return r
}
