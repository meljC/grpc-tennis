package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerGrpcHost   string `mapstructure:"server_grpc_host"`
	ServerGrpcPort   int    `mapstructure:"server_grpc_port"`
	DatabaseHost     string `mapstructure:"database_host"`
	DatabasePort     int    `mapstructure:"database_port"`
	DatabaseUser     string `mapstructure:"database_user"`
	DatabasePassword string `mapstructure:"database_password"`
	DatabaseName     string `mapstructure:"database_name"`
	DatabaseSslMode  string `mapstructure:"database_sslmode"`
	SeedUserCount    int    `mapstructure:"seed_user_count"`
}

type AppSettings struct {
	Config *Config
}

var App = &AppSettings{}

func Run() {
	var config = &Config{}

	viper.AddConfigPath("/Users/danielmeljanac/Projects/grpc-tennis")
	viper.SetConfigName("env_variables")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Error reading configuration file")
	}

	if err := viper.Unmarshal(config); err != nil {
		panic("Error unmarshalling config")
	}

	App.Config = config
	fmt.Println(App.Config)
}
