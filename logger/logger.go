// Only way for services to print logs.
package logger

import (
	"fmt"
	"go-base-end/operating_system"
	"go-base-end/resource"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	out *logrus.Logger
}

// singleton instance for current service
var global *Logger

// Default initializes default logger for service;
//
// This function depends on environment variable `IS_PROD_ENV`.
func Default() *Logger {
	// detect current environment
	producting := operating_system.GetEnv("IS_PROD_ENV", "false")

	logger := logrus.New()

	logger.SetOutput(os.Stderr)
	logger.SetLevel(logrus.InfoLevel)

	if producting == "true" {
		setJsonFormatter(logger)
	} else {
		setTextFormatter(logger)

		logger.ExitFunc = func(code int) {
			fmt.Printf("Trapped fatal log with return code %d", code)
		}
		logger.SetLevel(logrus.DebugLevel)
	}

	return &Logger{out: logger}
}

func (l Logger) Infof(format string, v ...any) {
	l.out.Infof(format, v...)
}

func (l Logger) Warnf(format string, v ...any) {
	l.out.Warnf(format, v...)
}

func (l Logger) Debugf(format string, v ...any) {
	l.out.Debugf(format, v...)
}

// Tracef is reserved but cannot print anything normally.
func (l Logger) Tracef(format string, v ...any) {
	l.out.Tracef(format, v...)
}

func (l Logger) Panicf(format string, v ...any) {
	l.out.Panicf(format, v...)
}

func (l Logger) Fatalf(format string, v ...any) {
	l.out.Fatalf(format, v...)
}

func (l *Logger) SetOutput(out io.Writer) {
	l.out.SetOutput(out)
}

func setJsonFormatter(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006/01/02 15:04:05.000",
	})
}

func setTextFormatter(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		DisableQuote:    true,
		TimestampFormat: "2006/01/02 15:04:05.000",
		FullTimestamp:   true,
	})
}

func init() {
	global = Default()
	resource.Register[*Logger](func(opts ...resource.Option) any {
		return global
	})
}
