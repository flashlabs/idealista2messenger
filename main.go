package main

import (
	"fmt"
	"github.com/flashlabs/idealista2messenger/internal/initializer"
	"github.com/flashlabs/idealista2messenger/internal/runner"
	"time"
)

const (
	configPath = "config"
)

func main() {
	fmt.Printf("Idealista2Messenger %s\n", time.Now().Format("2006.01.02 15:04:05"))

	config := initializer.InitConfig(configPath)
	initialize(config)
	runner.RunMainProcess(config)

	fmt.Printf("Application \"%s\" has finished processing\n", config.Application.Name)
}

func initialize(config *initializer.Config) {
	initializer.InitEnv(config)
}
