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

	configPath, fn := getInput()
	config := initializer.InitConfig(configPath)

	initialize(config)

	switch fn {
	case "runMainProcess":
		runner.RunMainProcess()
		break
	default:
		fmt.Println("Please provide function name. Possible options: runUpdateStock, runCreateProductSet, runClearAndDeleteProductSet, runIndexProductsProcess, runGetSimilarProductsProcess, runPreprocessData, runReportGenerator")
		break
	}

	fmt.Printf("Application \"%s\" has finished processing\n", config.Application.Name)
}

func getInput() (string, string) {
	cf := flag.String("config", "config", "Path to config file - relative to ./")
	fn := flag.String("fn", "runMainProcess", "Run given function")

	flag.Parse()

	return *cf, *fn
}

func initialize(config *initializer.Config) {
	// TODO
}
