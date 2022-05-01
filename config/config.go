package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	URL      string
	USERNAME string
	PASSWORD string
}

var NXRMConfig Config

func ReadConfig() Config {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("json")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file

	if err != nil { // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("config file not found %w", err))
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file %w ", err))
		}
	}
	NXRMConfig.URL = viper.GetString("url")
	NXRMConfig.USERNAME = viper.GetString("username")
	NXRMConfig.PASSWORD = viper.GetString("password")

	return NXRMConfig
}

func init() {
	NXRMConfig = ReadConfig()
}
