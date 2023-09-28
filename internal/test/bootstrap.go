package test

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const pathToRoot string = "../.."

func init() {
	testing.Init()
	flag.Parse()

	err := os.Chdir(rootPath())
	if err != nil {
		panic(err)
	}
}

func rootPath() string {
	_, currFile, _, _ := runtime.Caller(0)

	return filepath.Join(filepath.Dir(currFile), pathToRoot)
}
