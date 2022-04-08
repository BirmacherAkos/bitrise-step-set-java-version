package javasetter

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/log"
)

// JavaVersion ...
type JavaVersion string

// ...
const (
	JavaVersion8  = JavaVersion("8")
	JavaVersion11 = JavaVersion("11")
	JavaVersion17 = JavaVersion("17")
)

// Platform ...
type Platform string

// ...
const (
	MacOS  = Platform("MacOS")
	Ubuntu = Platform("Ubuntu")
)

func (j JavaSetter) platform() Platform {
	if runtime.GOOS == "linux" {
		return Ubuntu
	}
	return MacOS
}

type JavaSetter struct {
	logger     log.Logger
	cmdFactory command.Factory
}

func New(logger log.Logger, cmdFactory command.Factory) *JavaSetter {
	return &JavaSetter{logger: logger, cmdFactory: cmdFactory}
}

type Result struct {
	JavaHome string
}

func (j JavaSetter) SetJava(version JavaVersion) (Result, error) {
	j.logger.Println()
	j.logger.Infof("Checking platform")

	platform := j.platform()
	j.logger.Printf("Platform: %s", string(platform))

	j.logger.Println()
	j.logger.Infof("Running platform specific commands to set java version")
	switch platform {
	case MacOS:
		return j.setJavaMac(version)
	default:
		return j.setJavaUbuntu(version)
	}
}

func (j JavaSetter) setJavaMac(version JavaVersion) (Result, error) {
	if version == JavaVersion8 {
		version = "1.8"
	}

	//
	// jenv global
	cmdJenv := j.cmdFactory.Create(
		"jenv",
		[]string{"global", string(version)},
		&command.Opts{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		})

	j.logger.Printf("$ %s", cmdJenv.PrintableCommandArgs())
	if _, err := cmdJenv.RunAndReturnExitCode(); err != nil {
		return Result{}, err
	}

	//
	// jenv prefix
	cmdPrefix := j.cmdFactory.Create(
		"jenv",
		[]string{"prefix"},
		nil,
	)

	j.logger.Printf("$ %s", cmdPrefix.PrintableCommandArgs())
	javaHome, err := cmdPrefix.RunAndReturnTrimmedOutput()
	if err != nil {
		return Result{}, err
	}
	return Result{JavaHome: javaHome}, nil
}

func (j JavaSetter) setJavaUbuntu(version JavaVersion) (Result, error) {
	var javaHome string
	switch version {
	case JavaVersion8:
		javaHome = "/usr/lib/jvm/java-8-openjdk-amd64"
	case JavaVersion11:
		javaHome = "/usr/lib/jvm/java-11-openjdk-amd64"
	case JavaVersion17:
		javaHome = "/usr/lib/jvm/java-17-openjdk-amd64"
	}

	javaPath := filepath.Join(javaHome, "bin/java")
	javacPath := filepath.Join(javaHome, "bin/javac")
	javadocPath := filepath.Join(javaHome, "bin/javadoc")

	//
	// update-alternatives javac
	cmd := j.cmdFactory.Create(
		"sudo",
		[]string{
			"update-alternatives",
			"--set",
			"javac",
			javacPath,
		},
		&command.Opts{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		},
	)

	j.logger.Printf("$ %s", cmd.PrintableCommandArgs())
	if _, err := cmd.RunAndReturnExitCode(); err != nil {
		return Result{}, err
	}

	//
	// update-alternatives java
	cmd = j.cmdFactory.Create(
		"sudo",
		[]string{
			"update-alternatives",
			"--set",
			"java",
			javaPath,
		},
		nil,
	)

	j.logger.Printf("$ %s", cmd.PrintableCommandArgs())
	if _, err := cmd.RunAndReturnExitCode(); err != nil {
		return Result{}, err
	}

	//
	// update-alternatives javadoc
	cmd = j.cmdFactory.Create(
		"sudo",
		[]string{
			"update-alternatives",
			"--set",
			"javadoc",
			javadocPath,
		},
		nil,
	)

	j.logger.Printf("$ %s", cmd.PrintableCommandArgs())
	if _, err := cmd.RunAndReturnExitCode(); err != nil {
		return Result{}, err
	}

	return Result{JavaHome: javaHome}, nil
}
