package initializer

import (
	"github.com/joho/godotenv"
	"os"
)

func InitEnv(config *Config) {
	env := os.Getenv("I2M_ENV")
	if "" == env {
		env = "development"
	}

	_ = godotenv.Load(".env." + env + ".local")
	if "test" != env {
		_ = godotenv.Load(".env.local")
	}
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load() // The Original .env
}
