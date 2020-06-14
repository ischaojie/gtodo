package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/shiniao/gtodo/pkg/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"io"
	"os"
	"strings"
	"time"
)

const (
	WriterStdout = "stdout" // 标准输出
	WriterFile   = "file"   // 文件输出
)

const (
	RotateTimeDaily  = "daily" // 日志按天切割
	RotateTimeHourly = "houly" // 日志按小时切割
)

// zapLogger 对sugaredLogger的封装
type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

// NewZapLogger returns a new zap Logger
func NewZapLogger(cfg *Config) (Logger, error) {
	var cores []zapcore.Core
	var options []zap.Option
	encoder := getJSONEncoder()

	// 设置初始化字段: ip/app
	option := zap.Fields(zap.String("ip", util.GetLocalIP()), zap.String("app", viper.GetString("name")))
	options = append(options, option)

	// 获取日志写入类型writers
	writers := strings.Split(cfg.Writers, "',")

	allLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= zapcore.FatalLevel
	})

	for _, w := range writers {
		// 标准输出
		if w == WriterStdout {
			core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
			cores = append(cores, core)
		}
		// 写入文件
		if w == WriterFile {
			// 注意：如果多个文件，最后一个会是全的，前两个可能会丢日志
			infoFilename := cfg.LoggerFile
			infoWrite := getLogWriterWithTime(infoFilename)
			warnFilename := cfg.LoggerWarnFile
			warnWrite := getLogWriterWithTime(warnFilename)
			errorFilename := cfg.LoggerErrorFile
			errorWrite := getLogWriterWithTime(errorFilename)

			infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl <= zapcore.InfoLevel
			})
			warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl == zapcore.WarnLevel
			})
			errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.ErrorLevel
			})

			core := zapcore.NewCore(encoder, zapcore.AddSync(infoWrite), infoLevel)
			cores = append(cores, core)
			core = zapcore.NewCore(encoder, zapcore.AddSync(warnWrite), warnLevel)
			cores = append(cores, core)
			core = zapcore.NewCore(encoder, zapcore.AddSync(errorWrite), errorLevel)
			cores = append(cores, core)
		}
		if w != WriterFile && w != WriterStdout {
			core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
			cores = append(cores, core)
			allWriter := getLogWriterWithTime(cfg.LoggerFile)
			core = zapcore.NewCore(encoder, zapcore.AddSync(allWriter), allLevel)
			cores = append(cores, core)
		}
	}

	combinedCore := zapcore.NewTee(cores...)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	options = append(options, caller)
	// 开启文件和行号
	development := zap.Development()
	options = append(options, development)

	// 跳过文件调用层数
	addCallerSkip := zap.AddCallerSkip(2)
	options = append(options, addCallerSkip)

	// 构造新的logger
	logger := zap.New(combinedCore, options...).Sugar()
	return &zapLogger{sugaredLogger: logger}, nil
}

// getJSONEncoder return new JSONEncoder
func getJSONEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		NameKey:        "app",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		// EncodeDuration: zapcore.MillisDurationEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriterWithTime 按时间(小时)进行切割
func getLogWriterWithTime(filename string) io.Writer {
	logFullPath := filename
	hook, err := rotatelogs.New(
		logFullPath+".%Y%m%d%H",                                                // 时间格式使用shell的date时间格式
		rotatelogs.WithLinkName(logFullPath),                                   // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(viper.GetUint("logger.log_backup_count")), // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Hour),                                 // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// Debug logger
func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

// Info logger
func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

// Warn logger
func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

// Error logger
func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}
// Fatal logger
func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}
