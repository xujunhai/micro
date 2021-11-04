package logger

var (
	DefaultLogger, _ = NewZapLogger()
)

func SetLogger(logger Logger) {
	DefaultLogger = logger
}
