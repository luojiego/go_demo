package main

import (
	"go.uber.org/zap"
	"time"
)

func development() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("can't get url",
		zap.String("url", "https://luojiego.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
}

func production() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Debug("production debug can't get url",
		zap.String("url", "https://luojiego.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

	logger.Info("production info can't get url",
		zap.String("url", "https://luojiego.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
}

func main() {
	development()
	production()
}
