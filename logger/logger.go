// Package log provides a log interface
package logger

import (
	"gitlab.ziroom.com/rent-web/micro/logger/core"
	"os"
)

// Logger is a generic logging interface
type Logger interface {
	// Init initialises options
	Init(options ...core.Option) error
	// The Logger options
	Options() core.Options
	// Log writes a log entry
	Log(level core.Level, v ...interface{})
	// Logf writes a formatted log entry
	Logf(level core.Level, format string, v ...interface{})
	// Fields set fields to always be logged !! Deprecated
	Fields(fields map[string]interface{}) Logger

	Debug(msg string, fields ...core.Field)
	Info(msg string, fields ...core.Field)
	Warn(msg string, fields ...core.Field)
	Error(msg string, fields ...core.Field)
	Panic(msg string, fields ...core.Field)
	Fatal(msg string, fields ...core.Field)
	// String returns the name of logger
	String() string
}

func Init(opts ...core.Option) error {
	return DefaultLogger.Init(opts...)
}

func Fields(fields map[string]interface{}) Logger {
	return DefaultLogger.Fields(fields)
}

func Log(level core.Level, v ...interface{}) {
	DefaultLogger.Log(level, v...)
}

func Logf(level core.Level, format string, v ...interface{}) {
	DefaultLogger.Logf(level, format, v...)
}

func String() string {
	return DefaultLogger.String()
}

func Info(args ...interface{}) {
	DefaultLogger.Log(core.InfoLevel, args...)
}

func Infof(template string, args ...interface{}) {
	DefaultLogger.Logf(core.InfoLevel, template, args...)
}

func Trace(args ...interface{}) {
	DefaultLogger.Log(core.TraceLevel, args...)
}

func Tracef(template string, args ...interface{}) {
	DefaultLogger.Logf(core.TraceLevel, template, args...)
}

func Debug(args ...interface{}) {
	DefaultLogger.Log(core.DebugLevel, args...)
}

func Debugf(template string, args ...interface{}) {
	DefaultLogger.Logf(core.DebugLevel, template, args...)
}

func Warn(args ...interface{}) {
	DefaultLogger.Log(core.WarnLevel, args...)
}

func Warnf(template string, args ...interface{}) {
	DefaultLogger.Logf(core.WarnLevel, template, args...)
}

func Error(args ...interface{}) {
	DefaultLogger.Log(core.ErrorLevel, args...)
}

func Errorf(template string, args ...interface{}) {
	DefaultLogger.Logf(core.ErrorLevel, template, args...)
}

func Fatal(args ...interface{}) {
	DefaultLogger.Log(core.FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(template string, args ...interface{}) {
	DefaultLogger.Logf(core.FatalLevel, template, args...)
	os.Exit(1)
}

// Returns true if the given level is at or lower the current logger level
func V(lvl core.Level, logger Logger) bool {
	l := DefaultLogger
	if logger != nil {
		l = logger
	}
	return l.Options().Level <= lvl
}
