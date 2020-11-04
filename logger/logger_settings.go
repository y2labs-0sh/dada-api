package logger

import (
	"io"
	"os"
	"path"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

// Global log init setting
// 1. Log print to std.io & file
// 2. Keep 7days log,one log file a day
func init() {

	logDir := "./"

	logfile := path.Join(logDir, "dadaAPI")
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

	log.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// log.SetOutput(ansicolor.NewAnsiColorWriter(multiWriter))
	log.SetOutput(multiWriter)
	log.SetLevel(log.InfoLevel)
}

func Info(msg interface{}) func(...interface{}) {
	return log.WithFields(log.Fields{"Msg": msg}).Info
}

func Error(msg interface{}) func(...interface{}) {
	return log.WithFields(log.Fields{"Msg": msg}).Error
}

func Warning(msg interface{}) func(...interface{}) {
	return log.WithFields(log.Fields{"Msg": msg}).Warning
}

func Fatal(msg interface{}) func(...interface{}) {
	return log.WithFields(log.Fields{"Msg": msg}).Fatal
}
