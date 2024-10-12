package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	App      App
	Database Database
	Redis    Redis
}

type App struct {
	Port int `yaml:"port"`
}

type Database struct {
	Driver   string `yaml:"driver"`
	Url      string `yaml:"url"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
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
