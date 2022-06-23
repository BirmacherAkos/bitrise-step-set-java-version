package javasetter

import (
	"errors"
	"testing"

	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-set-java-version/javasetter/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_GivenMacShortestVersionAvailable_WhenSet_ThenIsSuccess(t *testing.T) {
	// Given
	logger := log.NewLogger()
	cmdFactory := new(mocks.CommandFactory)
	j := New(logger, cmdFactory)

	const javaHome = "/java11"

	cmdFactory.On("Create", "jenv", []string{"global", "11"}, mock.Anything).Once().Return(jenvGlobalSuccessCommand())
	cmdFactory.On("Create", "jenv", []string{"prefix"}, mock.Anything).Once().Return(jenvPrefixCommand(javaHome))

	// When
	got, err := j.setJavaMac(JavaVersion11)

	//Then
	require.NoError(t, err)
	require.Equal(t, Result{JavaHome: javaHome}, got)
	cmdFactory.AssertExpectations(t)
}

func Test_GivenMacNoVersionAvailable_WhenSet_ThenIsFailure(t *testing.T) {
	// Given
	logger := log.NewLogger()
	cmdFactory := new(mocks.CommandFactory)
	j := New(logger, cmdFactory)

	cmdFactory.On("Create", "jenv", []string{"global", "11"}, mock.Anything).Once().Return(jenvGlobalFailureCommand())
	cmdFactory.On("Create", "jenv", []string{"global", "11.0"}, mock.Anything).Once().Return(jenvGlobalFailureCommand())

	// When
	got, err := j.setJavaMac(JavaVersion11)

	//Then
	require.Error(t, err)
	require.Equal(t, Result{}, got)
	cmdFactory.AssertExpectations(t)
}

func Test_GivenMacFallbackVersionAvailable_WhenSet_ThenIsSuccess(t *testing.T) {
	// Given
	logger := log.NewLogger()
	cmdFactory := new(mocks.CommandFactory)
	j := New(logger, cmdFactory)

	const javaHome = "/java11"

	cmdFactory.On("Create", "jenv", []string{"global", "11"}, mock.Anything).Once().Return(jenvGlobalFailureCommand())
	cmdFactory.On("Create", "jenv", []string{"global", "11.0"}, mock.Anything).Once().Return(jenvGlobalSuccessCommand())
	cmdFactory.On("Create", "jenv", []string{"prefix"}, mock.Anything).Once().Return(jenvPrefixCommand(javaHome))

	// When
	got, err := j.setJavaMac(JavaVersion11)

	//Then
	require.NoError(t, err)
	require.Equal(t, Result{JavaHome: javaHome}, got)
	cmdFactory.AssertExpectations(t)
}

func jenvGlobalSuccessCommand() *mocks.Command {
	command := new(mocks.Command)

	command.On("PrintableCommandArgs").Return("")
	command.On("RunAndReturnTrimmedCombinedOutput").Return("", nil)

	return command
}

func jenvGlobalFailureCommand() *mocks.Command {
	command := new(mocks.Command)

	command.On("PrintableCommandArgs").Return("")
	command.On("RunAndReturnTrimmedCombinedOutput").Return("", errors.New("test_error"))

	return command
}

func jenvPrefixCommand(javaHome string) *mocks.Command {
	command := new(mocks.Command)

	command.On("PrintableCommandArgs").Return("")
	command.On("RunAndReturnTrimmedCombinedOutput").Return(javaHome, nil)

	return command
}
