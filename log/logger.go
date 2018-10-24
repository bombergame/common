package log

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: logrus.StandardLogger(),
	}
}

func (log *Logger) AsLogrusLogger() *logrus.Logger {
	return log.logger
}
