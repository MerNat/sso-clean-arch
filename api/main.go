package api

import (
	"context"
	"fmt"

	"github.com/mernat/sso-clean-arch/api/rest"
	"github.com/mernat/sso-clean-arch/config"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
)

func Run() {
	var err error
	tx := apm.DefaultTracer.StartTransaction("Init", "startup")
	defer tx.End()
	ctx := context.Background()
	ctx = apm.ContextWithTransaction(ctx, tx)
	globalEnvFileName := fmt.Sprintf("%s%s", "./envs/", "config.json")
	_config, _ := config.InitConfiguration(globalEnvFileName)

	logger := config.InitLogger(ctx, _config)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error("Failed to synchronize logger", zap.NamedError("error.message", err))
		}
	}(logger)

	logger = config.GetLoggerFromContext(ctx)
	logger = logger.With(apmzap.TraceContext(ctx)...)

	logger.Info(fmt.Sprintf("Starting service in %s environment", *_config.Environment))
	
	err = rest.ServeAPI(*_config.RestfulEndpoint)
	if err != nil {
		panic(err)
	}
}
