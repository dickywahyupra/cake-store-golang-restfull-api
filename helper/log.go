package helper

import "github.com/sirupsen/logrus"

func Log(msg interface{}) {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.Info(msg)
}
