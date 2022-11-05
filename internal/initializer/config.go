package initializer

import (
	"github.com/flashlabs/idealista2messenger/internal/structs"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
)

const appConfigFile = "app.yml"

type Config struct {
	Application structs.Application
}

func InitConfig(configPath string) *Config {
	config := &Config{}
	readCfg(&config.Application, configPath, appConfigFile)

	return config
}

func readCfg(config interface{}, configPath, configFile string) {
	path := filepath.Join(configPath, configFile)
	err := readYaml(config, path)
	if err != nil {
		log.Fatalf("failed to read config '%s' file: %s\n", path, err.Error())
	}
}

func readYaml(cfg interface{}, path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return err
	}

	defer func() {
		if ce := f.Close(); ce != nil {
			err = ce
		}
	}()

	return yaml.NewDecoder(f).Decode(cfg)
}
