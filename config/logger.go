package config

import (
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (appConfig *AppConfig) SetupLogger() {
	if strings.TrimSpace(appConfig.LoggerConfig.Filename) != "" {
		file, err := os.OpenFile(appConfig.LoggerConfig.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		appConfig.LoggerConfig.File = file
	}
	var writerSyncer zapcore.WriteSyncer
	if appConfig.LoggerConfig.File != nil {
		writerSyncer = zapcore.AddSync(appConfig.LoggerConfig.File)
	} else {
		writerSyncer = zapcore.AddSync(os.Stdout)
	}
	encoder := getEncoder(appConfig)
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.InfoLevel)
	appConfig.LoggerConfig.ZapLogger = zap.New(core, zap.AddCaller())
}

func getEncoder(appConf *AppConfig) zapcore.Encoder {
	if appConf.Environment == "dev" {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.FunctionKey = ""
		encoderConfig.CallerKey = ""
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
