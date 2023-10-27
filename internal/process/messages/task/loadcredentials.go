package task

import (
	"fmt"
	"os"
)

// Credentials reads creadentials from file and returns it.
func Credentials(file string) ([]byte, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error while reading credentials file: %w", err)
	}

	return b, nil
}
