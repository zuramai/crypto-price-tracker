package config

import (
	"database/sql"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppConfig struct {
	DB     *sql.DB
	Logger *zap.Logger
	Viper  *viper.Viper
}

func NewApp() *AppConfig {
	v := NewViper()
	l := NewLogger(v)
	db := NewDatabase(v, l)

	return &AppConfig{
		DB:     db,
		Logger: l,
		Viper:  v,
	}
}
