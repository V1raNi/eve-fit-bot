package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbPort           string `mapstructure:"DB_PORT"`
	DbHost           string `mapstructure:"DB_HOST"`
	DbName           string `mapstructure:"DB_NAME"`
	DbUser           string `mapstructure:"DB_USER"`
	DbPassword       string `mapstructure:"DB_PASSWORD"`
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
