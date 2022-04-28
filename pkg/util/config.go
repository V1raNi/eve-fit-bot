package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"PORT"`
	ConnectionString string `mapstructure:"CONNECTION"`
	Token            string `mapstructure:"TELEGRAM_TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	//config.Token = viper.GetString("TELEGRAM_TOKEN")
	return
}
