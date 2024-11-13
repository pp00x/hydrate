package util

import (
    "github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func SetupLogger() {
    Log = logrus.New()
    Log.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
}