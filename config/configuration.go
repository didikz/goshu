package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port     int
	Database string
}

var (
	configuration *Configuration
	once          sync.Once
)

func Load(configPath string) *Configuration {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configPath)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("unable to read config, %v", err)
		}

		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	})

	return configuration
}
