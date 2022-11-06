package task

import (
	"os"
)

func LoadCredentials() ([]byte, error) {
	b, err := os.ReadFile("config/credentials.json")
	if err != nil {
		return nil, err
	}

	return b, nil
}
