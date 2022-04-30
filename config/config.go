package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ReadConfig() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("json")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file %w ", err))
	}
	// uname := viper.GetString("username")
	// log.Println("Tool : ", viper.GetString("tool"))
	// log.Println("Url :", viper.GetString("url"))
	// log.Println("UserName :", viper.GetString("username"))
	// log.Println("Password :", viper.GetString("password"))

	toolConfing := viper.AllSettings()
	log.Println(toolConfing["url"])
}
