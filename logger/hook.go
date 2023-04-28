package logger

import "github.com/sirupsen/logrus"

type LogHook struct {
}

func (h *LogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *LogHook) Fire(entry *logrus.Entry) error {
	if entry.Context == nil {
		return nil
	}
	id := entry.Context.Value("log_id")
	if id == nil {
		return nil
	}
	entry.Data["log_id"] = id
	return nil
}
