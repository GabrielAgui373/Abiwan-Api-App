package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/gabrielagui373/obiwanapp-api/internal/config"
	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
	err      error
)

func InitDB(config *config.DBConfig) (*gorm.DB, error) {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=UTC", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

		instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			err = fmt.Errorf("failed to connect database: %w", err)
			return
		}

		err = instance.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Token{}, &models.Subject{}, models.Topic{})
		if err != nil {
			err = fmt.Errorf("failed to migrate database: %w", err)
			return
		}

		sqlDB, err := instance.DB()
		if err != nil {
			err = fmt.Errorf("failed to get undelying DB pool: %w", err)
			return
		}

		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetMaxOpenConns(50)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})

	return instance, nil
}
