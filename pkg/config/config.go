package config

import (
	logging "github.com/F1zm0n/bookey/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" env-default:"8080"`
	} `yaml:"server"`
	Database struct {
		Postgres struct {
			Driver   string `yaml:"driver"`
			User     string `yaml:"user"`
			Name     string `yaml:"name"`
			Password string `yaml:"password"`
			Sslmode  string `yaml:"sslmode"`
		} `yaml:"postgres"`
	} `yaml:"db"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("./config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
