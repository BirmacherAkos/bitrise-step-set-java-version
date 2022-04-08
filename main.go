package main

import (
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-steputils/v2/stepenv"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := log.NewLogger()
	envRepository := stepenv.NewRepository(env.NewRepository())
	cmdFactory := command.NewFactory(envRepository)
	inputParser := stepconf.NewInputParser(envRepository)

	javaSelector := NewJavaSelector(inputParser, envRepository, logger, cmdFactory)

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

	if err := javaSelector.Export(result); err != nil {
		logger.Errorf(err.Error())
		return 1
	}

	return 0
}
