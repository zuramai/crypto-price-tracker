package config

import (
	"database/sql"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewDatabase(viper *viper.Viper, log *zap.SugaredLogger) *sql.DB {
	connectionString := viper.GetString("database.connection_string")
	conn, err := sql.Open("sqlite", connectionString)

	if err != nil {
		log.Fatal("Error opening database connection", zap.Error(err))
	}

	return conn
}
