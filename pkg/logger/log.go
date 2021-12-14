package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogger() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true

	logrus.SetFormatter(customFormatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}
