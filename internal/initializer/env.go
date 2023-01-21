package initializer

import (
	"github.com/joho/godotenv"
	"os"
)

func Env() {
	env := os.Getenv("I2M_ENV")
	if env == "" {
		env = "development"
	}

	_ = godotenv.Load(".env." + env + ".local")
	if env != "test" {
		_ = godotenv.Load(".env.local")
	}
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load() // The Original .env
}
