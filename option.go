package logrustrace

import "github.com/sirupsen/logrus"

type Option func(*LogrusTrace)

func WithLevels(levels []logrus.Level) Option {
	return func(t *LogrusTrace) {
		t.levels = levels
	}
}

func WithTraceIDKey(key string) Option {
	return func(t *LogrusTrace) {
		t.traceIDKey = key
	}
}
