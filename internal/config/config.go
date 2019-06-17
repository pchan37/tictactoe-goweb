package config

import (
	"os"

	"github.com/PGo-Projects/output"
	"github.com/spf13/viper"
)

var Filename string
var DevRun bool

func Init() {
	if Filename != "" {
		viper.SetConfigFile(Filename)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		output.Error(err)
		os.Exit(1)
	}
}
