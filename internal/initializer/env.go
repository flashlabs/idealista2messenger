package initializer

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env() {
	env := os.Getenv("I2M_ENV")
	if env == "" {
		env = "development"
	}

	path := ".env." + env + ".local"
	err := godotenv.Load(path)

	if err != nil {
		log.Printf("error while loading .env file: %s", path)
	}

	if env != "test" {
		path = ".env.local"
		err = godotenv.Load(path)

		if err != nil {
			log.Printf("error while loading .env file: %s", path)
		}
	}

	path = ".env." + env
	err = godotenv.Load(path)

	if err != nil {
		log.Printf("error while loading .env file: %s", path)
	}

	err = godotenv.Load() // The Original .env
	if err != nil {
		log.Printf("error while loading default .env file")
	}
}
