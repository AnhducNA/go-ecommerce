package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 12))
	logger.Error("Error log", zap.Int("line", 15))

}

// format logs
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 1742990149.254945e+09 => 2025-03-26 18:55:55.195939377 +0700 +07
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encoderConfig.TimeKey = "time"
	// info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// caller -> main.go:12
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)

}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/error.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
