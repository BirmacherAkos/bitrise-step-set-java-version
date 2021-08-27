package javaSetter

import (
	"os"
	"runtime"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/log"
)

type JavaVersion string

const (
	JavaVersion_1_8 = JavaVersion("1.8")
	JavaVersion_11  = JavaVersion("11")
)

type Platform string

const (
	MacOS  = Platform("MacOS")
	Ubuntu = Platform("Ubuntu")
)

func (j JavaSetter) platform() Platform {
	if runtime.GOOS == "linux" {
		j.logger.Printf("Platform: Ubuntu")
		return Ubuntu
	}
	j.logger.Printf("Platform: MacOS")
	return MacOS
}

type JavaSetter struct {
	logger     log.Logger
	cmdFactory command.Factory
}

func New(logger log.Logger, cmdFactory command.Factory) *JavaSetter {
	return &JavaSetter{logger: logger, cmdFactory: cmdFactory}
}

func (j JavaSetter) SetJava(version JavaVersion) error {
	switch j.platform() {
	case MacOS:
		return j.setJavaMac(version)
	default:
		return j.setJavaUbuntu(version)
	}
}

func (j JavaSetter) setJavaMac(version JavaVersion) error {
	cmd_jenv := j.cmdFactory.Create(
		"jenv",
		[]string{"global", string(version)},
		&command.Opts{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		})
	j.logger.Println()
	j.logger.Printf("$ %s", cmd_jenv.PrintableCommandArgs())

	if _, err := cmd_jenv.RunAndReturnExitCode(); err != nil {
		return err
	}

	cmd_prefix := j.cmdFactory.Create(
		"$(jenv prefix)",
		[]string{},
		&command.Opts{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		})
	j.logger.Printf("$ %s", cmd_prefix.PrintableCommandArgs())
	jenvPrefix, err := cmd_prefix.RunAndReturnTrimmedOutput()

	if err != nil {
		return err
	}

	cmd_envman := j.cmdFactory.Create(
		"envman",
		[]string{"add", "--key", "JAVA_HOME", "--value", jenvPrefix},
		&command.Opts{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		})
	j.logger.Printf("$ %s", cmd_envman.PrintableCommandArgs())

	_, err = cmd_envman.RunAndReturnExitCode()
	return err
}

func (j JavaSetter) setJavaUbuntu(version JavaVersion) error {
	j.logger.Printf("sudo update-alternatives --set javac /usr/lib/jvm/java-8-openjdk-amd64/bin/javac...")
	return nil
}
