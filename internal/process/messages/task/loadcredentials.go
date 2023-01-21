package task

import (
	"os"
)

func Credentials(file string) ([]byte, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return b, nil
}
