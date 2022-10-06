package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type structure struct {
	Database struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"database"`
}

type Config struct {
	Db *gorm.DB
}

func GetConfig() (Config, error) {
	result := Config{}

	path, err := os.Getwd()
	if err != nil {
		return result, nil
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return result, err
	}

	config := structure{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return result, err
	}

	db, err := gorm.Open(mysql.Open(config.Database.Url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return result, err
	}

	result.Db = db

	return result, nil
}
