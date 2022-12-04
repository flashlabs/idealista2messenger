package task

import (
	"os"
)

func LoadCredentials(credentialsFile string) ([]byte, error) {
	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	return b, nil
}
