package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zdunker/gameStats"
	"github.com/zdunker/gameStats/config"
)

func main() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}
	gameStats.RunServer()
}

func loadConfig() error {
	viperConfig := viper.New()
	viperConfig.SetConfigName("config")
	viperConfig.AddConfigPath(".")
	viperConfig.SetConfigType("yaml")
	viperConfig.ReadInConfig()

	err := config.Load(viperConfig)
	if err != nil {
		return err
	}
	viperConfig.WatchConfig()

	viperConfig.OnConfigChange(func(e fsnotify.Event) {
		err = config.Load(viperConfig)
		if err != nil {
			//TODO: log error
		}
	})
	return nil
}
