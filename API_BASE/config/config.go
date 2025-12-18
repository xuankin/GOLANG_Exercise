package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DbSource           string        `mapstructure:"DB_SOURCE"`
	ServerAddr         string        `mapstructure:"SERVER_ADDRESS"`
	JWTSecret          string        `mapstructure:"JWT_SECRET"`
	JWTRefreshSecret   string        `mapstructure:"JWT_REFRESH_SECRET"`
	JWTDuration        time.Duration `mapstructure:"JWT_DURATION"`
	JWTRefreshDuration time.Duration `mapstructure:"JWT_REFRESH_DURATION"`
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
	return
}
