package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


type Logger struct{
	l *zap.Logger
}

func NewLogger(logFilePath string)*Logger{
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil{
		  fmt.Fprintf(os.Stderr, "cannot open log file: %v\n", err)
		  os.Exit(1)
	}

	defer logFile.Close()
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{logFilePath, "stdout"}
	config.ErrorOutputPaths = []string{logFilePath, "stderr"}

	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("02.01.2006 15:04:05.00000"))}

	logger, err := config.Build()
	if err != nil{
		log.Printf("Error via building logger: %v", err)
	}


	return &Logger{l:logger}
}


func (l *Logger) Err(msg string, fields ...zap.Field) {
    l.l.Error(msg, fields...)
}

func (l *Logger) Fatal(format string, v ...zap.Field){
	l.l.Error(format, v...)
}
func (l *Logger) Warn(format string, v ...zap.Field){
	l.l.Warn(format, v...)
}
func (l *Logger) Info(format string, v ...zap.Field){
	l.l.Info(format, v...)
}
