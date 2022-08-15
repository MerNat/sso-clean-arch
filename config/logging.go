package config

import (
	"context"

	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

type contextKey int

const (
	// Keys
	messageKey = "message"
	levelKey   = "log.level"
	timeKey    = "@timestamp"
	callerKey  = "caller"

	// Fields
	service                     = "service"
	name                        = "name"
	version                     = "version"
	tier                        = "tier"
	backend                     = "backend"
	environment                 = "environment"
	// ErrorMessage                = "error.message"


	ContextKeyLogger contextKey = iota
)

func InitLogger(ctx context.Context, config *ServiceConfig) *zap.Logger {
	serviceFields := map[string]interface{}{
		name:        config.Name,
		version:     config.Version,
		environment: config.Environment,
		tier:        backend,
	}

	initialFields := map[string]interface{}{
		service: serviceFields,
	}

	stdout := []string{"stdout"}
	stderr := []string{"stderr"}

	cfg := zap.Config{
		Level: zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   messageKey,
			LevelKey:     levelKey,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      timeKey,
			EncodeTime:   zapcore.RFC3339NanoTimeEncoder,
			CallerKey:    callerKey,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      stdout,
		ErrorOutputPaths: stderr,
		InitialFields:    initialFields,
	}

	// cfg := zap.NewProductionConfig()

	logger, err := cfg.Build(zap.WrapCore((&apmzap.Core{}).WrapCore))
	logger = logger.With(apmzap.TraceContext(ctx)...)

	if err != nil {
		panic(err)
	}

	logger.Info("Logger successfully initialized")
	Logger = logger
	return Logger
}

func getRootLogger() *zap.Logger {
	return Logger
}

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return getRootLogger()
	}

	value := ctx.Value(ContextKeyLogger)
	if value == nil {
		return getRootLogger()
	}

	return value.(*zap.Logger)
}
