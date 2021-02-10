package main

import (
	"context"
	"fmt"
	"github.com/autom8ter/eventgate/internal/helpers"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/server"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	configPath string
)

func init() {
	pflag.CommandLine.StringVar(&configPath, "config", helpers.EnvOr("EVENTGATE_CONFIG", "config.yaml"), "path to config file (env: EVENTGATE_CONFIG)")
	pflag.Parse()
}

func main() {
	c, err := server.ConfigFromFile(configPath)
	if err != nil {
		fmt.Printf("failed to read config file: %s", err.Error())
		return
	}
	var lgger = logger.New(
		c.Logging.Debug,
		zap.String("channel_provider", c.Backend.ChannelProvider.Name),
		zap.String("storage_provider", c.Backend.StorageProvider.Name),
	)
	lgger.Debug("loaded config", zap.Any("config", c))
	if err := server.ListenAndServe(context.Background(), lgger, c); err != nil {
		lgger.Error("eventgate server failure", zap.Error(err))
	}
}
