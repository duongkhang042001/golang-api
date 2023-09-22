package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	Logger = InitLogger()
}

func InitLogger() *zap.SugaredLogger {

	writerSyncer := getLogWriter()

	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())

	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		TimeKey:      "time",
		LevelKey:     "level",
		CallerKey:    "caller",
		EncodeLevel:  CustomLevelEncoder,
		EncodeTime:   SyslogTimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	})
}

func getLogWriter() zapcore.WriteSyncer {
	logsFileURL := fmt.Sprintf("./logs/logs_%s.log", time.Now().Format("2006_01_02"))
	file, _ := os.Create(logsFileURL)
	return zapcore.AddSync(file)
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func Info(msg string, args ...interface{}) {
	Logger.Infof(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	Logger.Debugf(msg, args...)
}

func Error(msg string, args ...interface{}) {
	Logger.Errorf(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	Logger.Warnf(msg, args...)
}

func Trace(msg string, args ...interface{}) {
	Logger.Debugf(msg, args...)
}
