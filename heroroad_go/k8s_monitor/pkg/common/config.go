package common

import (
	"fmt"

	"github.com/ory/viper"
	"github.com/spf13/pflag"
)

var (
	config GlobalConfig
)

type GlobalConfig struct {
	ServerPort int
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	pflag.IntVar(&config.ServerPort, "server-port", 8080, "Server port to listen on")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}
}

func GetGlobalConfig() *GlobalConfig {
	return &config
}
