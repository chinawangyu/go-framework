package logger

import (
	"fmt"
)

// Debug fmt.Sprintf to log a templated message.
func Debug(args ...interface{}) {
	Business.Logger.Debug(args...)
}

// Info uses fmt.Sprintf to log a templated message.
func Info(args ...interface{}) {
	Business.Logger.Info(args...)
}

// Warn uses fmt.Sprintf to log a templated message.
func Warn(args ...interface{}) {
	SendMonitor2DingDing(Business.dingUrl, args)
	Business.Logger.Warn(args...)
}

// Error uses fmt.Sprintf to log a templated message.
func Error(args ...interface{}) {
	SendMonitor2DingDing(Business.dingUrl, args)
	Business.Logger.Error(args...)
}

//// Fatal uses fmt.Sprintf to log a templated message.
//func Fatal(args ...interface{}) {
//	SendMonitor2DingDing(args)
//	SugaredLogger.Fatal(args...)
//}

// Debugf fmt.Sprintf to log a templated message.
func Debugf(format string, args ...interface{}) {
	Business.Logger.Debugf(format, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(format string, args ...interface{}) {
	Business.Logger.Infof(format, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	SendMonitor2DingDing(Business.dingUrl, str)
	Business.Logger.Warnf(format, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	SendMonitor2DingDing(Business.dingUrl, str)
	Business.Logger.Errorf(format, args...)
}
