package log

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

// Level shows logging verbosity
type Level int

const (
	// PanicLevel shows the highest severity
	// the log causes panic()
	PanicLevel Level = iota
	// FatalLevel log causes os.Exit(1)
	FatalLevel
	// ErrorLevel shows the error is caused
	ErrorLevel
	// WarnLevel shows it's non-critical, but will cause some unexpected habits
	WarnLevel
	// InfoLevel shows some general information
	// this level may default logging level
	InfoLevel
	// DebugLevel is the most verbose log level
	// usually, only enabled when development/debugging
	DebugLevel
)

// Logger manages logging level, format, and some more
type Logger struct {
	Level Level
	dest  io.Writer
}

type option func(*Logger)

// NewLogger returns a new Logger
func NewLogger(opts ...option) *Logger {
	l := Logger{}
	for _, opt := range opts {
		opt(&l)
	}
	return &l
}

// LoggingLevel option
func LoggingLevel(lv Level) func(*Logger) {
	return func(l *Logger) {
		l.Level = lv
	}
}

// SetOutput set the Logger's output destination
func (l *Logger) SetOutput(w io.Writer) {
	l.dest = w
}

// Write to Logger's destination
func (l *Logger) Write(b []byte) (int, error) {
	return l.dest.Write(b)
}

// Output writes if the log level should be written
func (l *Logger) Output(lv Level, s string) (int, error) {
	if l.Level < lv {
		return
	}
	n, err := l.Write([]byte(s))

	switch lv {
	case PanicLevel:
		panic(nil)
	case FatalLevel:
		os.Exit(1)
	}
	return n, err
}

// Print calls l.Output to print to logger
func (l *Logger) Print(lv Level, v ...interface{}) (int, error) {
	return l.Output(lv, fmt.Sprint(v...))
}

// Println calls l.Output to print to logger with trailing newline
func (l *Logger) Println(lv Level, v ...interface{}) (int, error) {
	return l.Output(lv, fmt.Sprintln(v...))
}

// Printf calls l.Output to print to logger with format string
func (l *Logger) Printf(lv Level, format string, v ...interface{}) (int, error) {
	return l.Output(lv, fmt.Sprintf(format, v...))
}

// Panic calls l.Output with PanicLevel and causes panic()
func (l *Logger) Panic(v ...interface{}) (int, error) {
	return l.Output(PanicLevel, fmt.Sprint(v...))
}

// Panicln calls l.Output with PanicLevel with trailing newline and causes panic()
func (l *Logger) Panicln(v ...interface{}) (int, error) {
	return l.Output(PanicLevel, fmt.Sprintln(v...))
}

// Panicf calls l.Output with PanicLevel with format string and causes panic()
func (l *Logger) Panicf(format string, v ...interface{}) (int, error) {
	return l.Output(PanicLevel, fmt.Sprintf(format, v...))
}

// Fatal calls l.Output with FatalLevel and then calls os.Exit(1)
func (l *Logger) Fatal(v ...interface{}) (int, error) {
	return l.Output(FatalLevel, fmt.Sprint(v...))
}

// Fatalln calls l.Output with FatalLevel with trailing newline and then calls os.Exit(1)
func (l *Logger) Fatalln(v ...interface{}) (int, error) {
	return l.Output(FatalLevel, fmt.Sprintln(v...))
}

// Fatalf calls l.Output with FatalLevel with format string and then calls os.Exit(1)
func (l *Logger) Fatalf(format string, v ...interface{}) (int, error) {
	return l.Output(FatalLevel, fmt.Sprintf(format, v...))
}

// Error calls l.Output with ErrorLevel and returns new error
func (l *Logger) Error(v ...interface{}) (int, error) {
	n, err := l.Output(ErrorLevel, fmt.Sprint(v...))
	if err != nil {
		return n, errors.Wrap(err, fmt.Sprint(v...))
	}
	return n, errors.New(fmt.Sprint(v...))
}

// Errorln calls l.Output with ErrorLevel with trailing newline and then calls os.Exit(1)
func (l *Logger) Errorln(v ...interface{}) (int, error) {
	n, err := l.Output(ErrorLevel, fmt.Sprintln(v...))
	if err != nil {
		return n, errors.Wrap(err, fmt.Sprintln(v...))
	}
	return n, errors.New(fmt.Sprintln(v...))
}

// Errorf calls l.Output with ErrorLevel with format string and then calls os.Exit(1)
func (l *Logger) Errorf(format string, v ...interface{}) (int, error) {
	n, err := l.Output(ErrorLevel, fmt.Sprintf(format, v...))
	if err != nil {
		return n, errors.Wrap(err, fmt.Sprintf(format, v...))
	}
	return n, fmt.Errorf(format, v...)
}

// Warn calls l.Output with WarnLevel
func (l *Logger) Warn(v ...interface{}) (int, error) {
	return l.Output(WarnLevel, fmt.Sprint(v...))
}

// Warnln calls l.Output with WarnLevel with trailing newline
func (l *Logger) Warnln(v ...interface{}) (int, error) {
	return l.Output(WarnLevel, fmt.Sprintln(v...))
}

// Warnf calls l.Output with WarnLevel with format string
func (l *Logger) Warnf(format string, v ...interface{}) (int, error) {
	return l.Output(WarnLevel, fmt.Sprintf(format, v...))
}

// Info calls l.Output with InfoLevel
func (l *Logger) Info(v ...interface{}) (int, error) {
	return l.Output(InfoLevel, fmt.Sprint(v...))
}

// Infoln calls l.Output with InfoLevel with trailing newline
func (l *Logger) Infoln(v ...interface{}) (int, error) {
	return l.Output(InfoLevel, fmt.Sprintln(v...))
}

// Infof calls l.Output with InfoLevel with format string
func (l *Logger) Infof(format string, v ...interface{}) (int, error) {
	return l.Output(InfoLevel, fmt.Sprintf(format, v...))
}

// Debug calls l.Output with DebugLevel
func (l *Logger) Debug(v ...interface{}) (int, error) {
	return l.Output(DebugLevel, fmt.Sprint(v...))
}

// Debugln calls l.Output with DebugLevel with trailing newline
func (l *Logger) Debugln(v ...interface{}) (int, error) {
	return l.Output(DebugLevel, fmt.Sprintln(v...))
}

// Debugf calls l.Output with DebugLevel with format string
func (l *Logger) Debugf(format string, v ...interface{}) (int, error) {
	return l.Output(DebugLevel, fmt.Sprintf(format, v...))
}
