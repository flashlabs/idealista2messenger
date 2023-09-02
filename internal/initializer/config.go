package initializer

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/flashlabs/idealista2messenger/internal/structs"
)

const (
	appConfig       = "app.yml"
	googleConfig    = "google.yml"
	messengerConfig = "messenger.yml"
)

type Config struct {
	Application structs.Application
	Google      structs.Google
	Messenger   structs.Messenger
}

func Cfg(dir string) *Config {
	c := &Config{}
	readCfg(&c.Application, dir, appConfig)
	readCfg(&c.Google, dir, googleConfig)
	readCfg(&c.Messenger, dir, messengerConfig)

	return c
}

func readCfg(target interface{}, dir, file string) {
	p := filepath.Join(dir, file)
	err := readYaml(target, p)
	if err != nil {
		log.Fatalf("failed to read config '%s' file: %s\n", p, err.Error())
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
