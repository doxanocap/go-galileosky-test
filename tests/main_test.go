package tests

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"testing"
	"todo/internal/manager"
	"todo/pkg/httpServer"
	"todo/pkg/logger"
	"todo/pkg/postgres"
)

var TestApp *App

type App struct {
	Server  *httpServer.Server
	Manager *manager.Manager

	pool *pgxpool.Pool
}

func getRouter() *gin.Engine {
	return TestApp.Manager.Processor().REST().Handler().Engine()
}

func TestMain(m *testing.M) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error in config file: %v", err)
	}

	viper.AutomaticEnv()

	logger.Init(viper.GetString("ENV_MODE"), viper.GetBool("ZAP_JSON"))

	TestApp = &App{
		Server:  httpServer.New(),
		Manager: manager.InitManager(),
	}

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
	TestApp.Manager.SetPool(conn)

	//if err := TestApp.Server.Run(TestApp.Manager.Processor().REST().Handler().Engine()); err != nil {
	//	logger.Log.Fatal("failed to run REST: %v", zap.Error(err))
	//}

	logger.Log.Info("TESTS STARTED")
	code := m.Run()
	logger.Log.Info("TESTS COMPLETED")

	os.Exit(code)
}
