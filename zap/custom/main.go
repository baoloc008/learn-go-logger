package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"sync"
)

var (
	sugarLogger *zap.SugaredLogger
	doOnce      sync.Once
)

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	SimpleHttpGet("www.google.com")
	SimpleHttpGet("http://www.google.com")
}

func InitLogger() {
	// Run Code Once on First Load (Concurrency Safe)
	doOnce.Do(func() {
		writerSyncer := getLogWriter()
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
		logger := zap.New(core, zap.AddCaller())
		sugarLogger = logger.Sugar()
	})
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./zap/custom/test.log", // location of log file
		MaxSize:    10,                      // maximum size of log file in MBs, before it is rotated
		MaxBackups: 5,                       // maximum no. of old files to retain
		MaxAge:     30,                      // maximum number of days it will retain old files
		Compress:   false,                   // whether to compress/archive old files
	}
	return zapcore.AddSync(lumberJackLogger)
}

func SimpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
