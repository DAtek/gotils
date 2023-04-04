package gotils

import (
	"io"
	"log"
	"strings"
)

type LogLevel uint8

const (
	LogLevelDebug = LogLevel(iota)
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelPanic
)

const ConfigLogLevel = EnvConfig("LOG_LEVEL")

const ErrorLogLevelNotSet = Error("LOG_LEVEL_NOT_SET")

type Logger struct {
	outLogger *log.Logger
	errLogger *log.Logger
	level     LogLevel
}

func (l *Logger) Debug(msg string, values ...any) {
	if l.level > LogLevelDebug {
		return
	}
	l.outLogger.Printf("DEBUG   | "+msg, values...)
}

func (l *Logger) Info(msg string, values ...any) {
	if l.level > LogLevelInfo {
		return
	}
	l.outLogger.Printf("INFO    | "+msg, values...)
}

func (l *Logger) Warning(msg string, values ...any) {
	if l.level > LogLevelWarning {
		return
	}
	l.outLogger.Printf("WARNING | "+msg, values...)
}

func (l *Logger) Error(msg string, values ...any) {
	if l.level > LogLevelError {
		return
	}
	l.outLogger.Printf("ERROR   | "+msg, values...)
}

func (l *Logger) Panic(msg string, values ...any) {
	l.errLogger.Panicf("PANIC   | "+msg, values...)
}

func NewLogger(level LogLevel, out, err io.Writer) *Logger {
	return &Logger{
		outLogger: log.New(out, "", log.Ldate|log.Ltime|log.LUTC),
		errLogger: log.New(err, "", log.Ldate|log.Ltime|log.LUTC),
		level:     level,
	}
}

func LogLevelFromEnv() (LogLevel, error) {
	logLevelMap := map[string]LogLevel{
		"debug":   LogLevelDebug,
		"info":    LogLevelInfo,
		"warning": LogLevelWarning,
		"error":   LogLevelError,
		"panic":   LogLevelPanic,
	}

	setting := ConfigLogLevel.Load()
	setting = strings.ToLower(setting)
	level, ok := logLevelMap[setting]

	if !ok {
		return LogLevelDebug, ErrorLogLevelNotSet
	}

	return level, nil
}
