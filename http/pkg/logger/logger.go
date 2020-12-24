package logger

import (
	"go.uber.org/zap"
)

type LogConfig struct {
	LogLevel     string //info error debug warning
	LogFile      string
	IsDebug      int //0-not debug(output only file)  1-debug (output either file and stdout )
	ConsoleDebug bool
	MaxSize      int
	MaxBackups   int
	MaxAge       int
	Compress     bool
}


var BusinessLogger *zap.SugaredLogger
var AccessLogger *zap.SugaredLogger


func InitBusinessLogger(logConf LogConfig) {
	BusinessLogger, _ = InitLogger(&logConf)
	BusinessLogger.Infof("log config:%+v", logConf)
}

func InitAccessLogger(logConf LogConfig) {
	AccessLogger, _ = InitLogger(&logConf)
	BusinessLogger.Infof("access log config:%+v", logConf)
}