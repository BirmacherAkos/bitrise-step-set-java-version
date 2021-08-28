package javaSetter

import (
	"os"
	"runtime"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/log"
)

type JavaVersion string

const (
	JavaVersion_8  = JavaVersion("8")
	JavaVersion_11 = JavaVersion("11")
)

type Platform string

const (
	MacOS  = Platform("MacOS")
	Ubuntu = Platform("Ubuntu")
)

const (
	UbuntuJavaPath_8   = "/usr/lib/jvm/java-8-openjdk-amd64/jre/bin/java"
	UbuntuJavaCPath_8  = "/usr/lib/jvm/java-8-openjdk-amd64/bin/javac"
	UbuntuJavaHome_1_8 = "/usr/lib/jvm/java-8-openjdk-amd64"

	UbuntuJavaPath_11  = "/usr/lib/jvm/java-11-openjdk-amd64/bin/java"
	UbuntuJavaCPath_11 = "/usr/lib/jvm/java-11-openjdk-amd64/bin/javac"
	UbuntuJavaHome_11  = "/usr/lib/jvm/java-11-openjdk-amd64"
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

func (j JavaSetter) SetJava(version JavaVersion) error {
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

func (j JavaSetter) setJavaMac(version JavaVersion) error {
	if version == JavaVersion_8 {
		version = JavaVersion("1.8")
	}

	//
	// jenv global
	cmd_jenv := j.cmdFactory.Create(
		"jenv",
		[]string{"global", string(version)},
		&command.Opts{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		})

	j.logger.Printf("$ %s", cmd_jenv.PrintableCommandArgs())
	if _, err := cmd_jenv.RunAndReturnExitCode(); err != nil {
		return err
	}

	//
	// jenv prefix
	cmd_prefix := j.cmdFactory.Create(
		"jenv",
		[]string{"prefix"},
		nil,
	)

	j.logger.Printf("$ %s", cmd_prefix.PrintableCommandArgs())
	jenvPrefix, err := cmd_prefix.RunAndReturnTrimmedOutput()
	if err != nil {
		return err
	}

	//
	// envman add
	cmd_envman := j.cmdFactory.Create(
		"envman",
		[]string{"add", "--key", "JAVA_HOME", "--value", jenvPrefix},
		nil,
	)

	j.logger.Printf("$ %s", cmd_envman.PrintableCommandArgs())
	_, err = cmd_envman.RunAndReturnExitCode()
	return err
}

func (j JavaSetter) setJavaUbuntu(version JavaVersion) error {
	javaPath, javaCPath, javaHome := func() (string, string, string) {
		switch version {
		case JavaVersion_8:
			return UbuntuJavaPath_8, UbuntuJavaCPath_8, UbuntuJavaHome_1_8
		case JavaVersion_11:
			return UbuntuJavaPath_11, UbuntuJavaCPath_11, UbuntuJavaHome_11
		default:
			return "", "", ""
		}
	}()

	//
	// update-alternatives javac
	cmd := j.cmdFactory.Create(
		"sudo",
		[]string{
			"update-alternatives",
			"--set",
			"javac",
			string(javaCPath),
		},
		&command.Opts{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		},
	)

	j.logger.Printf("$ %s", cmd.PrintableCommandArgs())
	if _, err := cmd.RunAndReturnExitCode(); err != nil {
		return err
	}

	//
	// update-alternatives java
	cmd = j.cmdFactory.Create(
		"sudo",
		[]string{
			"update-alternatives",
			"--set",
			"java",
			string(javaPath),
		},
		nil,
	)

	j.logger.Printf("$ %s", cmd.PrintableCommandArgs())
	if _, err := cmd.RunAndReturnExitCode(); err != nil {
		return err
	}

	//
	// envman JAVA_HOME
	cmd = j.cmdFactory.Create(
		"envman",
		[]string{
			"add",
			"--key",
			"JAVA_HOME",
			"--value",
			javaHome,
			string(javaPath),
		},
		nil,
	)

	j.logger.Printf("$ %s", cmd.PrintableCommandArgs())
	if _, err := cmd.RunAndReturnExitCode(); err != nil {
		return err
	}

	return nil
}
