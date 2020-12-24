package logger

import (
	"log"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugaredLogger *zap.SugaredLogger

func InitLogger(conf *LogConfig) (*zap.SugaredLogger, error) {

	SugaredLogger, err := initLogger(conf.LogFile, conf.LogLevel, conf.ConsoleDebug, conf.MaxSize, conf.MaxBackups, conf.MaxAge, conf.Compress)
	if err != nil {
		return SugaredLogger, err
	}

	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	return SugaredLogger, nil
}

func ZnTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func initLogger(logFile string, logLevel string, consoleDebug bool, maxSize, maxBackups, maxAge int, compress bool) (*zap.SugaredLogger, error) {
	hook := lumberjack.Logger{
		Filename:   logFile,    // ⽇志⽂件路径
		MaxSize:    maxSize,    // megabytes
		MaxBackups: maxBackups, // 最多保留3个备份
		MaxAge:     maxAge,     //days
		Compress:   compress,   // 是否压缩 disabled by default
		LocalTime:  true,
	}
	fileWriter := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)

	// for human operators.
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = ZnTimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	//初始化所有core
	var allCore []zapcore.Core

	if consoleDebug {
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleDebugging, level))
	}

	allCore = append(allCore, zapcore.NewCore(consoleEncoder, fileWriter, level))

	core := zapcore.NewTee(allCore...)

	// From a zapcore.Core, it's easy to construct a Logger.
	zlog := zap.New(core).WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))

	SugaredLogger = zlog.Sugar()
	return SugaredLogger, nil
}
