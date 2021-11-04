package jaeger

import (
	"gitlab.ziroom.com/rent-web/micro/logger"
	"gitlab.ziroom.com/rent-web/micro/logger/core"
)

//adapter jaeger logger interface
var JaegerLogger = &jaegerLogger{Logger: logger.DefaultLogger}

type jaegerLogger struct {
	logger.Logger
}

func (j *jaegerLogger) Error(msg string) {
	j.Logger.Logf(core.ErrorLevel, "ERROR: %s", msg)
}

// Infof logs a message at info priority
func (j *jaegerLogger) Infof(msg string, args ...interface{}) {
	j.Logger.Logf(core.InfoLevel, msg, args...)
}
