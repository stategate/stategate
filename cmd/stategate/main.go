package main

import (
	"context"
	"fmt"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/autom8ter/stategate/internal/server"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func init() {
	godotenv.Load()
}

func main() {
	c := &server.Config{}
	if err := c.LoadEnv(); err != nil {
		fmt.Printf("failed to read environmental variables: %s", err.Error())
		return
	}
	c.SetDefaults()
	if err := c.Validate(); err != nil {
		fmt.Printf("failed to validate config: %s", err.Error())
		return
	}
	var lgger = logger.New(
		c.Debug,
		zap.Any("cache_provider", cast.ToString(c.CacheProvider["name"])),
		zap.Any("storage_provider", cast.ToString(c.StorageProvider["name"])),
	)
	lgger.Debug("loaded config from env", zap.Any("config", c))
	if err := server.ListenAndServe(context.Background(), lgger, c); err != nil {
		lgger.Error("stategate server failure", zap.Error(err))
	}
}
