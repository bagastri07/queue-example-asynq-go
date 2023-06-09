package logger

import "github.com/sirupsen/logrus"

func SetUpLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}
