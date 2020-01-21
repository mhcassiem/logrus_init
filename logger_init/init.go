package logger_init

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func Init_logger() *logrus.Logger{
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   false,
			ForceFormatting: true,
		},
	}
	return log
}
