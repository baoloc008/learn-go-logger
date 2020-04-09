package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func test() {
	InitLogger()
	defer logger.Sync()
	SimpleHttpGet("www.google.com")
	SimpleHttpGet("http://www.google.com")
}

func InitLogger() {
	logger = zap.NewExample()
}

func SimpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
