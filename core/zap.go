package core

import (
	"fmt"
	"os"
	"path/filepath"
	"test-case-gin/utils"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.SugaredLogger) {
	logMode := zapcore.DebugLevel
	if utils.Config.AppMode == "debug" {
		logMode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriter(), zapcore.AddSync(os.Stdout)), logMode)

	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriter() zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format("2006-01-02") + ".txt"
	fmt.Println(stLogFilePath)

	lumberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    utils.ZapConfig.MaxSize,    // megabytes
		MaxBackups: utils.ZapConfig.MaxBackups, // files
		MaxAge:     utils.ZapConfig.MaxAge,     // days
		Compress:   utils.ZapConfig.Compress,   // disabled by default
	}
	return zapcore.AddSync(lumberjackSyncer)
}
