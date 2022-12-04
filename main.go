package main

import (
	"flag"
	"fmt"
	"github.com/flashlabs/idealista2messenger/internal/initializer"
	"github.com/flashlabs/idealista2messenger/internal/runner"
	"time"
)

func main() {
	fmt.Printf("Idealista2Messenger %s\n", time.Now().Format("2006.01.02 15:04:05"))

	configPath := getInput()
	config := initializer.InitConfig(configPath)

	initialize(config)

	runner.RunMainProcess(config)

	fmt.Printf("Application \"%s\" has finished processing\n", config.Application.Name)
}

func getInput() string {
	cf := flag.String("config", "config", "Path to config file - relative to ./")

	flag.Parse()

	return *cf
}

func initialize(config *initializer.Config) {
	initializer.InitEnv(config)
}
