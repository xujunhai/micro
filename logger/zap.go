package logger

import (
	"context"
	"xmicro/common/constant"
	"xmicro/logger/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type ZapLog struct {
	c     *zap.Logger
	lv    *zap.AtomicLevel
	opts  core.Options
	sugar *zap.SugaredLogger
}

// New builds a new logger based on options
func NewZapLogger(opts ...core.Option) (Logger, error) {
	// Default options
	options := core.Options{
		Level:  core.InfoLevel,
		Fields: make(map[string]interface{}),
		//Out:     os.Stdout,
		CallerSkipCount: 2, //wrapper caller 2
		Context:         context.Background(),
	}

	l := &ZapLog{opts: options}
	if err := l.Init(opts...); err != nil {
		return nil, err
	}

	return l, nil
}

func (l *ZapLog) Sugar() *zap.SugaredLogger {
	return l.sugar
}

// Init initialises options
func (l *ZapLog) Init(options ...core.Option) error {
	for _, o := range options {
		o(&l.opts)
	}

	var logFile string
	// log path
	if file, ok := l.opts.Context.Value(Key{}).(string); ok {
		logFile = file
	} else if e := os.Getenv(constant.AppLogConfFile); e != "" {
		logFile = e
	} else {
		logFile = l.opts.File
	}

	writeSyn := getLogWriter(logFile)
	encoder := getEncoder()

	var level = loggerToZapLevel(l.opts.Level)
	core := zapcore.NewCore(encoder, writeSyn, level)
	l.c = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(l.opts.CallerSkipCount))

	//还原sugar caller
	c := *l.c
	copy := c.WithOptions(zap.AddCallerSkip(1))
	l.sugar = copy.Sugar()
	return nil
}

func loggerToZapLevel(level core.Level) zapcore.Level {
	switch level {
	case core.TraceLevel, core.DebugLevel:
		return zap.DebugLevel
	case core.InfoLevel:
		return zap.InfoLevel
	case core.WarnLevel:
		return zap.WarnLevel
	case core.ErrorLevel:
		return zap.ErrorLevel
	case core.FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func getLogWriter(logFile string) zapcore.WriteSyncer {
	if len(logFile) == 0 {
		return zapcore.AddSync(os.Stdout)
	}
	//加入日志分割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1000,
		MaxBackups: 20,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	zapConf := zap.NewProductionEncoderConfig()
	zapConf.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewJSONEncoder(zapConf)
}

// The Logger options
func (l *ZapLog) Options() core.Options {
	return l.opts
}

// Log writes a log entry
func (l *ZapLog) Log(level core.Level, v ...interface{}) {
	zapLevel := loggerToZapLevel(level)
	l.sugarLog(zapLevel, "", v...)
}

func (l *ZapLog) sugarLog(zapLevel zapcore.Level, format string, v ...interface{}) {
	isFormat := len(format) > 0
	switch zapLevel {
	case zap.DebugLevel:
		if isFormat {
			l.sugar.Debugf(format, v...)
		} else {
			l.sugar.Debug(v...)
		}
	case zap.InfoLevel:
		if isFormat {
			l.sugar.Infof(format, v...)
		} else {
			l.sugar.Info(v...)
		}
	case zap.FatalLevel:
		if isFormat {
			l.sugar.Fatalf(format, v...)
		} else {
			l.sugar.Fatal(v)
		}
	case zap.WarnLevel:
		if isFormat {
			l.sugar.Warnf(format, v...)
		} else {
			l.sugar.Warn(v...)
		}
	case zap.ErrorLevel:
		if isFormat {
			l.sugar.Errorf(format, v...)
		} else {
			l.sugar.Error(v...)
		}
	}
}

// Logf writes a formatted log entry
func (l *ZapLog) Logf(level core.Level, format string, v ...interface{}) {
	zapLevel := loggerToZapLevel(level)
	l.sugarLog(zapLevel, format, v...)
}

// Inline opt
func loggerToZapField(fields []core.Field) []zap.Field {
	var zapFields []zap.Field
	for _, v := range fields {
		zapFields = append(zapFields, zap.Any(v.Key, v.Value))
	}

	return zapFields
}

// String returns the name of logger
func (l *ZapLog) Debug(msg string, fields ...core.Field) {
	zapFields := loggerToZapField(fields)
	l.c.Debug(msg, zapFields...)
}

func (l *ZapLog) Info(msg string, fields ...core.Field) {
	zapFields := loggerToZapField(fields)
	l.c.Info(msg, zapFields...)
}

func (l *ZapLog) Warn(msg string, fields ...core.Field) {
	zapFields := loggerToZapField(fields)
	l.c.Warn(msg, zapFields...)
}

func (l *ZapLog) Error(msg string, fields ...core.Field) {
	zapFields := loggerToZapField(fields)
	l.c.Error(msg, zapFields...)
}

func (l *ZapLog) Panic(msg string, fields ...core.Field) {
	zapFields := loggerToZapField(fields)
	l.c.Panic(msg, zapFields...)
}

func (l *ZapLog) Fatal(msg string, fields ...core.Field) {
	zapFields := loggerToZapField(fields)
	l.c.Fatal(msg, zapFields...)
}

// Fields set fields to always be logged
func (l *ZapLog) Fields(fields map[string]interface{}) Logger {
	return l
}

func (l *ZapLog) String() string {
	return "zap"
}
