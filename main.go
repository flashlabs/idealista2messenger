package main

import (
	"fmt"
	"time"

	"github.com/flashlabs/idealista2messenger/internal/initializer"
	"github.com/flashlabs/idealista2messenger/internal/runner"
)

const (
	configDirPath = "config"
)

func main() {
	fmt.Printf("Idealista2Messenger %s\n", time.Now().Format("2006.01.02 15:04:05"))

	config := initializer.Cfg(configDirPath)
	initializer.Env()
	runner.MainProcess(config)

	fmt.Printf("Application \"%s\" has finished processing\n", config.Application.Name)
}
