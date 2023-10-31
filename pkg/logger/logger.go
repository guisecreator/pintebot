package logger

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func InitLogger() *logrus.Logger {
	log := logrus.New()
	log.Formatter = new(prefixed.TextFormatter)
	log.Level = logrus.DebugLevel
	return log
}
