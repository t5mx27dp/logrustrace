package logrustrace

import (
	"github.com/sirupsen/logrus"
	"github.com/t5mx27dp/trace"
)

type LogrusTrace struct {
	levels []logrus.Level

	traceIDKey string
}

var _ logrus.Hook = (*LogrusTrace)(nil)

func New(opts ...Option) *LogrusTrace {
	t := &LogrusTrace{
		levels:     logrus.AllLevels,
		traceIDKey: "TraceID",
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

func (t *LogrusTrace) Levels() []logrus.Level {
	return t.levels
}

func (t *LogrusTrace) Fire(entry *logrus.Entry) error {
	entry.Data[t.traceIDKey] = trace.TraceID(entry.Context)
	return nil
}
