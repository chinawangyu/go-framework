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
	DingUrl      string
}

type loggerBase struct {
	Logger  *zap.SugaredLogger
	dingUrl string
}

var Business *loggerBase

//初始化业务日志
func InitBusinessLogger(logConf *LogConfig) error {
	loggerObj, err := InitLogger(logConf)
	if err != nil {
		return err
	}

	Business = &loggerBase{
		Logger:  loggerObj,
		dingUrl: logConf.DingUrl,
	}
	Business.Logger.Infof("InitBusinessLogger config:%+v", logConf)
	return nil
}
