package driver

import (
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

// LogDriver ...
type LogDriver struct {
	logFName string
	level    logrus.Level
}

// NewLogDriver ...
func NewLogDriver(filename string, level logrus.Level) *LogDriver {
	return &LogDriver{
		logFName: filename,
		level:    level,
	}
}

// InitLog ...
func (l *LogDriver) InitLog() {
	dt := time.Now()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logs/" + dt.Format("20060102") + "_" + l.logFName,
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Level:      l.level,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	logrus.SetLevel(l.level)
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	logrus.AddHook(rotateFileHook)
}
