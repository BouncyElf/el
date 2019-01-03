package el

import (
	"io"
)

type Conf struct {
	Prefix      string
	Outs        []io.Writer
	Value       map[string]interface{}
	Format      string
	AddCaller   bool
	CallerSkip  int
	AddTime     bool
	TimeFormat  string
	NotPanic    bool
	NotFatal    bool
	LowestLevel level
}

func SetConf(c *Conf) {
	theLogger.SetConf(c)
}

func (l *Logger) SetConf(c *Conf) {
	l.c = c
}
