package el

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := New()
	assert.EqualValues(t, DefaultConf(), l.c)
}

func TestLog(t *testing.T) {
	c := DefaultConf()
	c.NotFatal = true
	c.NotPanic = true
	l := New(c)
	m := Map{
		"meaningless map": true,
	}
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
	l.Panic("panic")
	l.Fatal("fatal")
	l.Debug("debug", m)
	l.Info("info", m)
	l.Warn("warn", m)
	l.Error("error", m)
	l.Panic("panic", m)
	l.Fatal("fatal", m)

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
