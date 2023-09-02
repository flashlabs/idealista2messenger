package test

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const suffixPath = "idealista2messenger"

func init() {
	testing.Init()
	flag.Parse()

	err := os.Chdir(getRootPath())
	if err != nil {
		panic(err)
	}
}

func getRootPath() string {
	wd, _ := os.Getwd()
	for flag.Lookup("test.v") != nil && !strings.HasSuffix(wd, suffixPath) {
		wd = filepath.Dir(wd)
	}

	return wd
}
