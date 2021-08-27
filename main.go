package main

import (
	"os"

	"github.com/bitrise-io/go-steputils/stepconf"
	"github.com/bitrise-io/go-steputils/stepenv"
	"github.com/bitrise-io/go-utils/env"
	"github.com/bitrise-io/go-utils/log"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := log.NewLogger()
	envRepository := stepenv.NewRepository(env.NewRepository())
	inputParser := stepconf.NewInputParser(envRepository)

	javaSelector := NewJavaSelector(inputParser, envRepository, logger)

	config, err := javaSelector.ProcessConfig()
	if err != nil {
		logger.Errorf(err.Error())
		return 1
	}

	result, err := javaSelector.Run(config)
	if err != nil {
		logger.Errorf(err.Error())
		return 1
	}

	logger.Printf("Version: %s", result.version)

	if err := javaSelector.Export(result); err != nil {
		logger.Errorf(err.Error())
		return 1
	}

	return 0
}
