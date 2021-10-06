package logger

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel log.Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

// var Fields = log.Fields
var logger *log.Logger

func SetLevel(level log.Level) {
	Logger().SetLevel(level)
}

func Logger() *log.Logger {
	if nil == logger {
		logger = log.New()
		logger.SetLevel(log.DebugLevel)
		logger.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
			// TimestampFormat: time.RFC3339,
			TimestampFormat: "2006-01-02 15:04:05.000",
		})
	}
	return logger
}

func WithFields(fields Fields) *log.Entry {
	return Logger().WithTime(time.Now()).WithFields(log.Fields(fields))
}

func WithField(key string, value interface{}) *log.Entry {
	return Logger().WithTime(time.Now()).WithField(key, value)
}

// func Debug(args ...interface{}) {
// 	Logger().WithFields(log.Fields{}).Debug(args...)
// }

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}
