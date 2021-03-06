package logger

import (
	env "github.com/fabiotavarespr/go-env"
	"sync"

	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option struct {
	ServiceName    string
	ServiceVersion string
	Environment    string
	LogLevel       string
}

var (
	once    sync.Once
	skipper = zap.AddCallerSkip(1)
	option  *Option
)

func Init(opt *Option) {
	once.Do(func() {
		option = opt
		zapLogger, err := newZap()

		if err != nil {
			panic(err)
		}

		zap.ReplaceGlobals(zapLogger)
	})
}

func Sync() error {
	return zap.L().Sync()
}

func Debug(message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Debug(message, makeAttributes(attr)...)
}

func Info(message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Info(message, makeAttributes(attr)...)
}

func Warn(message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Warn(message, makeAttributes(attr)...)
}

func Fatal(message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Fatal(message, makeAttributes(attr)...)
}

func Error(message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Error(message, makeAttributes(attr)...)
}

func newZap(options ...zap.Option) (*zap.Logger, error) {
	return newZapConfig().Build(options...)
}

func newZapConfig() zap.Config {
	logLevel := zap.NewAtomicLevel()
	err := logLevel.UnmarshalText([]byte(option.LogLevel))

	if err != nil {
		panic(err)
	}

	log := zap.Config{
		Level:       logLevel,
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    newZapEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if env.GetEnvWithDefaultAsString("ENVIRONMENT", "prod") == "prod" {
		log.Development = false
	}

	return log
}

func newZapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "Timestamp",
		LevelKey:       "Severity",
		NameKey:        "Logger",
		CallerKey:      "Caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "Message",
		StacktraceKey:  "Stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func makeAttributes(attr attributes.Attributes) []zapcore.Field {
	if attr == nil {
		attr = attributes.New()
	}

	return []zapcore.Field{
		zap.Any("Attributes", attr),
		zap.Any("Resource", map[string]interface{}{
			"service.name":        option.ServiceName,
			"service.version":     option.ServiceVersion,
			"service.environment": option.Environment,
		}),
	}
}
