package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewLogger(viper *viper.Viper) *zap.Logger {
	var logger *zap.Logger

	if viper.GetString("env") == "development" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	return logger
}
