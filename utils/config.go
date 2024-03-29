package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	BotToken         string `mapstructure:"BOT_TOKEN"`
	BotVersion       string `mapstructure:"BOT_VERSION"`
	BotOwner         string `mapstructure:"BOT_OWNER"`
	BotChannelId     string `mapstructure:"BOT_CHANNEL_ID"`
	BotGuildId       string `mapstructure:"BOT_GUILD_ID"`
	BotApplicationId string `mapstructure:"BOT_APPLICATION_ID"`
	DatabaseUrl      string `mapstructure:"DATABASE_URL"`
}

func LoadConfig(configPath, configName string) (Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	log.Println("CONFIG SUCCESSFULLY LOADED")
	return config, nil
}
