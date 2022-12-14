package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

var (
	// Environment is the environment for the application
	Environment = os.Getenv("GIN_ENV")
)

// DefaultDBName is the default name for the database. Defaults to <app-name>-<environment>
var DefaultDBName string

func init() {
	if Environment == "" {
		Environment = "development"
	}
	DefaultDBName = dbPrefix + "-" + Environment

	if Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
		tmpdir := filepath.Join(os.TempDir(), "memorize")
		os.MkdirAll(tmpdir, 0700)
		logFileName := filepath.Join(tmpdir, fmt.Sprintf("%s-%d.log", Environment, time.Now().Unix()))
		logFile, err := os.Create(logFileName)

		if err == nil {
			logrus.SetOutput(logFile)
		} else {
			logrus.WithFields(logrus.Fields{"err": err.Error()}).Error("error while opening log file.")
		}
		logrus.SetFormatter(&logrus.TextFormatter{DisableSorting: true, DisableColors: true})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{DisableSorting: true})
	}
}

// vi:syntax=go

