package el

import (
	"testing"
)

func TestLog(t *testing.T) {
	c := DefaultConf()
	c.NotFatal = true
	c.NotPanic = true
	m := Map{
		"meaningless map": true,
	}
	SetConf(c)
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
	Panic("panic")
	Fatal("fatal")
	Debug("debug", m)
	Info("info", m)
	Warn("warn", m)
	Error("error", m)
	Panic("panic", m)
	Fatal("fatal", m)
}
