package utils

import "github.com/sirupsen/logrus"

// LogLevel represents the level of logging
type LogLevel uint8

const (
	// InfoLevel represents info level logs
	InfoLevel LogLevel = iota
	// WarningLevel represents warning level logs
	WarningLevel
	// ErrorLevel represents error level logs
	ErrorLevel
)

type Log struct {
	Level LogLevel
	Info  string
}

// Logger wraps the logrus logger
type Logger struct {
	log *logrus.Logger
}

// NewLogger creates and configures a new logger instance
func NewLogger() *Logger {
	l := logrus.New()

	// Configure logrus to use colored output
	l.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	})

	return &Logger{
		log: l,
	}
}

// LogMessage logs a message with the specified level and info
func (l *Logger) LogMessage(logEntry Log) {
	switch logEntry.Level {
	case InfoLevel:
		l.log.WithFields(logrus.Fields{
			"level": "info",
		}).Info(logEntry.Info)
	case WarningLevel:
		l.log.WithFields(logrus.Fields{
			"level": "warning",
		}).Warn(logEntry.Info)
	case ErrorLevel:
		l.log.WithFields(logrus.Fields{
			"level": "error",
		}).Error(logEntry.Info)
	default:
		l.log.WithFields(logrus.Fields{
			"level": "unknown",
		}).Info(logEntry.Info)
	}
}
