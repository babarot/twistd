package twistd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

type JSONFormatter struct{}

type logWrapper struct{}

func (f *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := entry.Data
	entry.Data = logrus.Fields{}

	entry.Data["data"] = data
	entry.Data["time"] = entry.Time.Format(time.RFC3339)
	entry.Data["level"] = strings.ToUpper(entry.Level.String())

	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marchal fields to JSON, %v", err)
	}

	return append(serialized, '\n'), nil
}

var Log = func() *logrus.Logger {
	log := logrus.New()

	log.Formatter = new(JSONFormatter)

	level, err := logrus.ParseLevel("info")
	if err != nil {
		panic(err)
	}
	log.Level = level

	return log
}()

var Logger = new(logWrapper)

func (l *logWrapper) Info(data map[string]interface{}) {
	l.log(data, func(e *logrus.Entry) { e.Info("") })
}

func (l *logWrapper) Warn(data map[string]interface{}) {
	l.log(data, func(e *logrus.Entry) { e.Warn("") })
}

func (l *logWrapper) Error(data map[string]interface{}) {
	l.log(data, func(e *logrus.Entry) { e.Error("") })
}

func (l *logWrapper) log(data map[string]interface{}, fn func(*logrus.Entry)) {
	logPath := "/Users/b4b4r07/src/github.com/b4b4r07/twistd/cmd/twistd/twistd.log"

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	Log.Out = f

	if err != nil {
		panic(err)
	}

	fn(Log.WithFields(data))
}
