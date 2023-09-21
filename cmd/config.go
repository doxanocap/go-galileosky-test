package cmd

import (
	"github.com/spf13/viper"
	"log"
)

var envConfigs = struct {
	PORT         string `mapstructure:"PORT"`
	EnvMode      string `mapstructure:"ENV_MODE"`
	ZapJSON      bool   `mapstructure:"ZAP_JSON"`
	PsqlHost     string `mapstructure:"PSQL_HOST"`
	PsqlPort     string `mapstructure:"PSQL_PORT"`
	PsqlUser     string `mapstructure:"PSQL_USER"`
	PsqlPassword string `mapstructure:"PSQL_PASSWORD"`
	PsqlDB       string `mapstructure:"PSQL_DB"`
	PsqlSSL      string `mapstructure:"PSQL_SSL"`
}{}

func InitConfig() {
	viper.SetDefault("ENV_MODE", "development")
	viper.SetDefault("ZAP_JSON", false)

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error in config file: %v \n", err)
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&envConfigs); err != nil {
		log.Printf("unable to unmarshal .env: %v \n", err)
	}
}
