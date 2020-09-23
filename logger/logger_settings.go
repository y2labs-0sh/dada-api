package logger

import (
	"io"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

// Global log init setting
// 1. Log print to std.io & file
// 2. Keep 7days log,one log file a day
func init() {

	logDir := "./"

	//_, err := os.Stat(logDir)
	//if os.IsExist(err) {
	//	if err := os.MkdirAll(logDir, 0755); err != nil {
	//		panic(err)
	//	}
	//}

	logfile := path.Join(logDir, "aggr")
	fsWriter, err := rotatelogs.New(
		logfile+"_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Duration(168)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(fsWriter, os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(multiWriter)
	log.SetLevel(log.InfoLevel)
}
