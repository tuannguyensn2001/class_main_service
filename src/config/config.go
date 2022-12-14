package config

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
)

type structure struct {
	Database struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"database"`
	App struct {
		Port      string `mapstructure:"port"`
		SecretKey string `mapstructure:"secretKey"`
	} `mapstructure:"app"`
	Jaeger string `mapstructure:"jaeger"`
	Redis  struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"redis"`
}

type Config struct {
	Db        *gorm.DB
	Port      string
	Jaeger    string
	SecretKey string
	Rd        *redis.Client
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
	viper.SetEnvPrefix(viper.GetString("ENV"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err = viper.ReadInConfig()

	bind := map[string]string{
		"database.url":  "DATABASE_URL",
		"app.port":      "PORT",
		"jaeger":        "JAEGER",
		"app.secretKey": "SECRET_KEY",
	}

	for key, val := range bind {
		err = viper.BindEnv(key, val)
		if err != nil {
			return result, err
		}
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

	rd := redis.NewClient(&redis.Options{
		Addr: config.Redis.Url,
	})

	ping := rd.Ping(context.Background())
	if ping.Err() != nil {
		zap.S().Error(err)
	}

	result.Db = db
	result.Port = config.App.Port
	result.Jaeger = config.Jaeger
	result.SecretKey = config.App.SecretKey
	result.Rd = rd

	return result, nil
}
