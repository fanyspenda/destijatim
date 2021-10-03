package main

import (
	"destijatim/configs"
	"destijatim/routers"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	configs.InitDB()
	routers.InitRoute()
}
