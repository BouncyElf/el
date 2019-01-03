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
	DebugL level = iota
	InfoL
	WarnL
	ErrorL
	PanicL
	FatalL
)

type level int
type Map map[string]interface{}

type Logger struct {
	c *Conf
}

var theLogger *Logger
var levelMap map[level]string

func init() {
	theLogger = new(Logger)
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

func New(c ...*Conf) *Logger {
	if len(c) == 0 {
		c = append(c, DefaultConf())
	}
	return &Logger{
		c: c[0],
	}
}

func DefaultConf() *Conf {
	return &Conf{
		Prefix:      "",
		Outs:        []io.Writer{os.Stdout},
		Value:       nil,
		Format:      "[%s] %s: %s. %+v\n",
		AddCaller:   true,
		CallerSkip:  2,
		AddTime:     true,
		TimeFormat:  "unixnano",
		NotPanic:    false,
		NotFatal:    false,
		LowestLevel: InfoL,
	}
}

func Debug(msg string, m map[string]interface{}) {
	theLogger.log(DebugL, msg, m)
}

func Info(msg string, m map[string]interface{}) {
	theLogger.log(InfoL, msg, m)
}

func Warn(msg string, m map[string]interface{}) {
	theLogger.log(WarnL, msg, m)
}

func Error(msg string, m map[string]interface{}) {
	theLogger.log(ErrorL, msg, m)
}

func Panic(msg string, m map[string]interface{}) {
	theLogger.log(PanicL, msg, m)
	if theLogger.c.NotPanic {
		return
	}
	panic(msg)
}

func Fatal(msg string, m map[string]interface{}) {
	theLogger.log(FatalL, msg, m)
	if theLogger.c.NotFatal {
		return
	}
	os.Exit(1)
}

func (l *Logger) Debug(msg string, m map[string]interface{}) {
	l.log(DebugL, msg, m)
}

func (l *Logger) Info(msg string, m map[string]interface{}) {
	l.log(InfoL, msg, m)
}

func (l *Logger) Warn(msg string, m map[string]interface{}) {
	l.log(WarnL, msg, m)
}

func (l *Logger) Error(msg string, m map[string]interface{}) {
	l.log(ErrorL, msg, m)
}

func (l *Logger) Panic(msg string, m map[string]interface{}) {
	l.log(PanicL, msg, m)
	if l.c.NotPanic {
		return
	}
	panic(msg)
}

func (l *Logger) Fatal(msg string, m map[string]interface{}) {
	l.log(FatalL, msg, m)
	if l.c.NotFatal {
		return
	}
	os.Exit(1)
}

func (l *Logger) log(ll level, msg string, m map[string]interface{}) {
	if ll < l.c.LowestLevel {
		return
	}
	if m == nil {
		m = make(map[string]interface{})
	}
	if l.c.AddCaller {
		_, file, line, _ := runtime.Caller(l.c.CallerSkip)
		m["caller"] = fmt.Sprintf("file:%s, line: %d", file, line)
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
	finalLogMsg := fmt.Sprintf(l.c.Format, levelMap[ll], l.c.Prefix, msg, m)
	for _, v := range l.c.Outs {
		v.Write([]byte(finalLogMsg))
	}
}
