package main

import (
	"github.com/fabiotavarespr/go-env"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/topics"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/platform/consumers"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger.Init(&logger.Option{
		ServiceName:    env.GetEnvWithDefaultAsString("WORKER_NAME", "worker"),
		ServiceVersion: env.GetEnvWithDefaultAsString("WORKER_VERSION", "0.0.1"),
		Environment:    env.GetEnvWithDefaultAsString("ENVIRONMENT", "dev"),
		LogLevel:       env.GetEnvWithDefaultAsString("LOG_LEVEL", "info"),
	})

	defer logger.Sync()

	startTime := time.Now()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		details := attributes.New().WithError(errors.New("interrupt signal detected"))
		details["uptime"] = time.Since(startTime).String()
		details["signal"] = sig.Signal

		logger.Info("Starting worker", details)
		os.Exit(0)
	}()

	c := consumers.Consumer{
		Broker: env.GetEnvWithDefaultAsString("BROKER_ADDRESS", "localhost:9092"),
		Group:  env.GetEnvWithDefaultAsString("CONSUMER_GROUP", "test-consumer-group"),
		Topic:  topics.TopicOrderReceived,
	}

	log.Fatal(c.SubscribeAndListen())
}
