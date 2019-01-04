package el

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	// log level
	DebugL level = iota
	InfoL
	WarnL
	ErrorL
	PanicL
	FatalL
)

type level int
type Map map[string]interface{}

type logger struct {
	c *Conf
}

// theLogger is the Global logger.
var theLogger *logger

// levelMap map level to string.
var levelMap map[level]string

func init() {
	theLogger = new(logger)
	theLogger.c = DefaultConf()
	levelMap = map[level]string{
		DebugL: "debug",
		InfoL:  "info",
		WarnL:  "warn",
		ErrorL: "error",
		PanicL: "panic",
		FatalL: "fatal",
	}
}

// DefaultConf returns the default conf.
func DefaultConf() *Conf {
	return &Conf{
		Prefix:      "",
		Outs:        []io.Writer{os.Stdout},
		Value:       nil,
		Format:      "[%s]%s %s. %s\n",
		AddCaller:   true,
		CallerSkip:  2,
		AddTime:     true,
		TimeFormat:  "unixnano",
		NotPanic:    false,
		NotFatal:    false,
		LowestLevel: DebugL,
	}
}

// Debug is theLogger's debug method.
func Debug(msg string, m ...map[string]interface{}) {
	theLogger.log(DebugL, msg, m...)
}

// Info is theLogger's info method.
func Info(msg string, m ...map[string]interface{}) {
	theLogger.log(InfoL, msg, m...)
}

// Warn is theLogger's warn method.
func Warn(msg string, m ...map[string]interface{}) {
	theLogger.log(WarnL, msg, m...)
}

// Error is theLogger's error method.
func Error(msg string, m ...map[string]interface{}) {
	theLogger.log(ErrorL, msg, m...)
}

// Panic is theLogger's panic method.
func Panic(msg string, m ...map[string]interface{}) {
	theLogger.log(PanicL, msg, m...)
	if theLogger.c.NotPanic {
		return
	}
	panic(msg)
}

// Fatal is theLogger's fatal method.
func Fatal(msg string, m ...map[string]interface{}) {
	theLogger.log(FatalL, msg, m...)
	if theLogger.c.NotFatal {
		return
	}
	os.Exit(1)
}

// log is the logging method.
func (l *logger) log(ll level, msg string, ms ...map[string]interface{}) {
	if ll < l.c.LowestLevel {
		return
	}
	m := make(map[string]interface{})
	if len(ms) != 0 {
		for _, mm := range ms {
			for k, v := range mm {
				m[k] = v
			}
		}
	}
	if l.c.AddCaller {
		_, file, line, _ := runtime.Caller(l.c.CallerSkip)
		m["caller"] = fmt.Sprintf("%s:%d", file, line)
	}
	if l.c.AddTime {
		t := ""
		if l.c.TimeFormat == "unix" {
			t = strconv.FormatInt(time.Now().Unix(), 10)
		} else if l.c.TimeFormat == "unixnano" {
			t = strconv.FormatInt(time.Now().UnixNano(), 10)
		} else {
			t = time.Now().Format(l.c.TimeFormat)
		}
		m["time"] = t
	}
	for k, v := range l.c.Value {
		m[k] = v
	}
	mString := fmt.Sprintf("%+v", m)
	mString = mString[4:]
	mString = mString[:len(mString)-1]
	finalLogMsg := fmt.Sprintf(l.c.Format, levelMap[ll], l.c.Prefix, msg, mString)
	for _, v := range l.c.Outs {
		v.Write([]byte(finalLogMsg))
	}
}
