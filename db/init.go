package db

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jay-bhogayata/product-service/config"
	"github.com/jay-bhogayata/product-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) error {

	c := &config.AppCfg.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", c.Host, c.User, c.Password, c.DbName, c.Port, c.Sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

	slog.Info("database connected successfully...")

	db.AutoMigrate(&models.Category{})

	cfg.Db = db

	return nil
}
