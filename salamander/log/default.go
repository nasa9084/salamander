package log

var _defaultLogger *Logger

func init() {
	_defaultLogger = NewLogger(LoggingLevel(InfoLevel))
}
