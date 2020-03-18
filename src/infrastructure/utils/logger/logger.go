package logger

import (
	"github.com/sirupsen/logrus"
)

type LEVEL int

type ILogger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

const (
	TRACE LEVEL = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

func (l LEVEL) String() string {
	return [...]string{"trace", "debug", "info", "warn", "error", "fatal"}[l]
}
