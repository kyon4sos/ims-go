package util

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"img-server/core"
)

func ReadConfig() core.AppConfig {
	viper.SetConfigName("application")         // name of config file (without extension)
	viper.SetConfigType("json")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/")  // path to look for the config file in
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	var appConfig core.AppConfig
	fmt.Println(viper.AllSettings())
	err = mapstructure.Decode(viper.AllSettings(),&appConfig)
	if err != nil {
		panic(err.Error())
	}
	return appConfig
}
