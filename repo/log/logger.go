package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path"
)

type Level int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

var (
	opt    *Options
	logger *zap.Logger
)

func InitLog(o *Options) error {
	opt = o
	var err error
	logger, err = NewRollingLogger(opt.FilePath, opt.FileName)
	return err
}

// NewLogger 简单的logger，不会自动分割和滚动
func NewLogger(filePath string, fileName string) (*zap.Logger, error) {
	//日志路径
	logFile := path.Join(filePath, fileName)
	var zc = zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
			TimeKey:    "time",
			NameKey:    "name",
			CallerKey:  "caller",
			//FunctionKey:    "function",
			StacktraceKey:  "stacktrace",
			SkipLineEnding: false,
			LineEnding:     zapcore.DefaultLineEnding, //定义换行符
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder, //全路径编码
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths:      []string{"stdout", logFile},
		ErrorOutputPaths: []string{"stderr"},
	}
	return zc.Build()
}

// NewRollingLogger 会自动分割数据，info和error日志分开记录
func NewRollingLogger(filePath string, fileName string) (*zap.Logger, error) {
	var coreArr []zapcore.Core
	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
		TimeKey:    "time",
		NameKey:    "name",
		CallerKey:  "caller",
		//FunctionKey:    "function",
		StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding, //定义换行符
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, //全路径编码
		EncodeName:     zapcore.FullNameEncoder,
	}
	//copnsole 输出普通日志，jsonencoder输出json日志
	//encoder:=zapcore.NewConsoleEncoder(encoderConfig)
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	//日志级别
	logLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.Level(opt.LogLevel)

	})
	//日志路径
	logFile := path.Join(filePath, fileName)
	//文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,        //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    opt.MaxSize,    //文件大小限制，单位MB
		MaxAge:     opt.MaxAge,     //最大保留日志天数
		MaxBackups: opt.MaxBackups, //最大保留日志文件数量
		Compress:   false,          //是否压缩处理
	})
	//file core
	fileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer), logLevel)
	coreArr = append(coreArr, fileCore)

	//console core
	if opt.Console {
		consoleCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(ColorSink{}), zap.WarnLevel)
		coreArr = append(coreArr, consoleCore)
	}
	//zap.AddCaller()为显示文件名和行号，可省略
	logger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
	return logger, nil
}

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func Debugf(msg string, args ...interface{}) {
	logger.Sugar().Debugf(msg, args...)
}

func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

func Infof(msg string, args ...interface{}) {
	logger.Sugar().Infof(msg, args...)
}
func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

func Warnf(msg string, args ...interface{}) {
	logger.Sugar().Warnf(msg, args...)
}

func Error(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}

func Errorf(msg string, args ...interface{}) {
	logger.Sugar().Errorf(msg, args...)
}

func Panic(msg string, fields ...zapcore.Field) {
	logger.Panic(msg, fields...)
}

func Panicf(msg string, args ...interface{}) {
	logger.Sugar().Panicf(msg, args...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Fatalf(msg string, args ...interface{}) {
	logger.Sugar().Fatalf(msg, args...)
}
