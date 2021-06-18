package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DSN string `mapstructure:"dsn"`
}

func ReadConfig() Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}
	c := Config{}
	c.DSN = viper.GetString("database.dsn")
	return c
}
