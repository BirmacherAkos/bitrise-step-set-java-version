package main

import (
	"os"

	"github.com/BirmacherAkos/bitrise-step-set-java-version/javaSetter"
	"github.com/bitrise-io/go-steputils/stepconf"
	"github.com/bitrise-io/go-steputils/stepenv"
	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/env"
	"github.com/bitrise-io/go-utils/log"
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

	if err := javaSelector.Run(config); err != nil {
		logger.Errorf(err.Error())
		return 1
	}

	if err := javaSelector.Export(javaSetter.JavaVersion(config.javaVersion)); err != nil {
		logger.Errorf(err.Error())
		return 1
	}

	return 0
}
