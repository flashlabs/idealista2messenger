package runner

import (
	"fmt"
	"github.com/flashlabs/idealista2messenger/internal/process/mainprocess"
)

func RunMainProcess() bool {
	fmt.Println("Main process")

	success, err := mainprocess.Execute()
	if err != nil {
		fmt.Println(err)

		return false
	}

	return success
}
