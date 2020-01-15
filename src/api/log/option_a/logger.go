package option_a

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tv2169145/golang-microservices/src/api/config"
	"os"
	"strings"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level: level,
		Out: os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}
	//if config.IsProduction() {
	//	Log.Formatter = &logrus.JSONFormatter{}
	//} else {
	//	//Log.Formatter = &logrus.JSONFormatter{}
	//	Log.Formatter = &logrus.TextFormatter{}
	//}
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(ParseFields(tags...)).Info(msg)
}

func Error(msg string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %s", msg, err.Error())
	Log.WithFields(ParseFields(tags...)).Error(msg)
}

func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(ParseFields(tags...)).Debug(msg)
}

func ParseFields(tags ...string) logrus.Fields {
	//result := make(logrus.Fields, len(tags))
	result := logrus.Fields{}
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	return result
}
