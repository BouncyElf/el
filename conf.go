package el

import (
	"io"
)

type Conf struct {
	// Prefix is the prefix of a log.
	Prefix string
	// The logs are write to these outputs.
	Outs []io.Writer
	// Value is the value, that every log will include.
	Value map[string]interface{}
	// Format is the log format.
	Format string
	// AddCaller specific should the log add caller.
	AddCaller bool
	// CallerSkip specific the runtime.Caller(skip) skip number.
	CallerSkip int
	// AddTime specific should the log add time.
	AddTime bool
	// TimeFormat specific the log's time format.
	// "unixnano" and "unix" indicates time.UnixNano() and time.Unix()
	TimeFormat string
	// NotPanic specific should the log panic after panic method.
	NotPanic bool
	// NotFatal specific should the log exit after panic method.
	NotFatal bool
	// LowestLevel specific the lowest log level.
	// level less than LowestLevel will not be log.
	LowestLevel level
}

// SetConf set theLogger's conf.
func SetConf(c *Conf) {
	theLogger.SetConf(c)
}

// SetConf set the l's conf.
func (l *Logger) SetConf(c *Conf) {
	l.c = c
}
