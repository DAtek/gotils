package gotils

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	loggingScenarios := []struct {
		loggerMethod      string
		expectedLevelInfo string
		level             LogLevel
	}{
		{loggerMethod: "Debug", expectedLevelInfo: "DEBUG", level: LogLevelDebug},
		{loggerMethod: "Info", expectedLevelInfo: "INFO", level: LogLevelInfo},
		{loggerMethod: "Warning", expectedLevelInfo: "WARNING", level: LogLevelWarning},
		{loggerMethod: "Error", expectedLevelInfo: "ERROR", level: LogLevelError},
	}
	for _, scenario := range loggingScenarios {
		t.Run(fmt.Sprintf("logs %s message", scenario.loggerMethod), func(t *testing.T) {
			out := &bytes.Buffer{}
			err := &bytes.Buffer{}
			logger := NewLogger(scenario.level, out, err)
			loggerReflection := reflect.ValueOf(logger)
			method := loggerReflection.MethodByName(scenario.loggerMethod)
			msg := "hey"

			method.Call(
				[]reflect.Value{reflect.ValueOf(msg)},
			)

			assert.True(t, strings.Contains(out.String(), msg))
			assert.True(t, strings.Contains(out.String(), scenario.expectedLevelInfo))
			assert.Equal(t, "", err.String())
		})
	}

	noLoggingScenarios := []struct {
		loggerMethod string
		level        LogLevel
	}{
		{loggerMethod: "Debug", level: LogLevelInfo},
		{loggerMethod: "Info", level: LogLevelWarning},
		{loggerMethod: "Warning", level: LogLevelError},
		{loggerMethod: "Error", level: LogLevelPanic},
	}
	for _, scenario := range noLoggingScenarios {
		t.Run(fmt.Sprintf("logs %s message", scenario.loggerMethod), func(t *testing.T) {
			out := &bytes.Buffer{}
			err := &bytes.Buffer{}
			logger := NewLogger(scenario.level, out, err)
			loggerReflection := reflect.ValueOf(logger)
			method := loggerReflection.MethodByName(scenario.loggerMethod)
			msg := "hey"

			method.Call(
				[]reflect.Value{reflect.ValueOf(msg)},
			)

			assert.Equal(t, "", out.String())
			assert.Equal(t, "", err.String())
		})
	}

	t.Run("Logs panic message", func(t *testing.T) {
		out := &bytes.Buffer{}
		err := &bytes.Buffer{}

		logger := NewLogger(LogLevelDebug, out, err)
		msg := "hey"

		assert.Panics(t, func() {
			logger.Panic(msg)
		})

		assert.True(t, strings.Contains(err.String(), msg))
		assert.True(t, strings.Contains(err.String(), "PANIC"))
		assert.Equal(t, "", out.String())
	})
}

func TestLogLevelFromEnv(t *testing.T) {
	scenarios := []struct {
		logLevel      string
		expectedLevel LogLevel
	}{
		{logLevel: "debug", expectedLevel: LogLevelDebug},
		{logLevel: "DEBUG", expectedLevel: LogLevelDebug},
		{logLevel: "info", expectedLevel: LogLevelInfo},
		{logLevel: "warning", expectedLevel: LogLevelWarning},
		{logLevel: "error", expectedLevel: LogLevelError},
		{logLevel: "panic", expectedLevel: LogLevelPanic},
	}
	for _, scenario := range scenarios {
		t.Run(fmt.Sprintf("Loads log level from env - %s", scenario.logLevel), func(t *testing.T) {
			os.Setenv("LOG_LEVEL", scenario.logLevel)
			defer os.Unsetenv("LOG_LEVEL")

			level, err := LogLevelFromEnv()

			assert.Nil(t, err)
			assert.Equal(t, scenario.expectedLevel, level)
		})
	}

	t.Run("Returns error if log level not set", func(t *testing.T) {
		_, err := LogLevelFromEnv()

		assert.EqualError(t, err, ErrorLogLevelNotSet.Error())
	})

}
