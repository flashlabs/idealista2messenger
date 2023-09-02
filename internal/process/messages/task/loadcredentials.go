package task

import (
	"os"
)

// Credentials reads creadentials from file and returns it
func Credentials(file string) ([]byte, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return b, nil
}
