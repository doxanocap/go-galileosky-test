package cmd

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"todo/pkg/logger"
	"todo/pkg/postgres"
)

func (app *App) ConnectToPostgres() {
	pgDSN := postgres.GetDSN(postgres.Config{
		Host:     viper.GetString("PSQL_HOST"),
		Port:     viper.GetString("PSQL_PORT"),
		Username: viper.GetString("PSQL_USER"),
		Password: viper.GetString("PSQL_PASSWORD"),
		DBName:   viper.GetString("PSQL_DB"),
		SSLMode:  viper.GetString("PSQL_SSL"),
	})
	conn, err := postgres.Connect(context.Background(), pgDSN)

	if err != nil {
		logger.Log.Fatal("failed connection to postgres:", zap.Error(err))
	}
	app.Manager.SetPool(conn)
}
