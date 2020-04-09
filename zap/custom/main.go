package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
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
	//return os.Stdout // write to console
	file, _ := os.Create("./zap/custom/test.log")
	return zapcore.AddSync(file)
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
