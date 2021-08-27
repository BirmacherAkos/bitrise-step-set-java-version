package main

import (
	"fmt"

	"github.com/BirmacherAkos/bitrise-step-set-java-version/javaSetter"
	"github.com/bitrise-io/go-steputils/stepconf"
	"github.com/bitrise-io/go-steputils/tools"
	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/env"
	"github.com/bitrise-io/go-utils/log"
)

// Config is the Set java version step configuration
type Input struct {
	JavaVersion string `env:"set_java_version,opt[11,1.8]"`
}

type Config struct {
	javaVersion string
}

// JavaSelector ...
type JavaSelector struct {
	inputParser   stepconf.InputParser
	envRepository env.Repository
	logger        log.Logger
	cmdFactory    command.Factory
}

// NewActivateSSHKey ...
func NewJavaSelector(inputParser stepconf.InputParser, envRepository env.Repository, logger log.Logger, cmdFactory command.Factory) *JavaSelector {
	return &JavaSelector{inputParser: inputParser, envRepository: envRepository, logger: logger, cmdFactory: cmdFactory}
}

// ProcessConfig ...
func (j JavaSelector) ProcessConfig() (Config, error) {
	var input Input
	err := j.inputParser.Parse(&input)
	if err != nil {
		return Config{}, err
	}
	stepconf.Print(input) // TODO: log.Infof(stepconf.toString(input))
	return Config{
		javaVersion: input.JavaVersion,
	}, nil
}

// Run ...
func (j JavaSelector) Run(cfg Config) error {
	versionToSet := javaSetter.JavaVersion(cfg.javaVersion)
	setter := javaSetter.New(j.logger, j.cmdFactory)
	err := setter.SetJava(versionToSet)

	return err
}

// Export ...
func (j JavaSelector) Export(version javaSetter.JavaVersion) error {
	if string(version) == "" {
		return nil
	}
	if err := tools.ExportEnvironmentWithEnvman("JAVA_VERSION", string(version)); err != nil {
		return fmt.Errorf("failed to export environment variable: %s", "JAVA_VERSION")
	}
	return nil
}
