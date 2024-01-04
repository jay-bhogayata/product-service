package config

import (
	"log/slog"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DbName   string
		Sslmode  string
	}
	Db *gorm.DB
}

var AppCfg Config

func Init() error {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&AppCfg)
	if err != nil {
		return err
	}

	slog.Info("config loaded successfully...")

	return nil
}
