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

var log = func() *logrus.Logger {
	log := logrus.New()

	log.Formatter = new(JSONFormatter)

	level, err := logrus.ParseLevel("info")
	if err != nil {
		panic(err)
	}
	log.Level = level

	return log
}()

func (t *Twistd) Info(data map[string]interface{}) {
	t.log(data, func(e *logrus.Entry) { e.Info("") })
}

func (t *Twistd) Warn(data map[string]interface{}) {
	t.log(data, func(e *logrus.Entry) { e.Warn("") })
}

func (t *Twistd) Error(data map[string]interface{}) {
	t.log(data, func(e *logrus.Entry) { e.Error("") })
}

func (t *Twistd) log(data map[string]interface{}, fn func(*logrus.Entry)) {
	var conf ConfToml
	if err := LoadConf(t.Option.Config, &conf); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(conf.Core.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	log.Out = f

	if err != nil {
		panic(err)
	}

	fn(log.WithFields(data))
}
