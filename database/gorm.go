package database

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"golang-project-layout/config"
)

func NewDefaultConfig() *config.GormConfig {
	return &config.GormConfig{
		Config:             newGormConfig(),
		MaxIdleConnections: 2,
		MaxOpenConnections: 4,
	}
}

func newGormConfig() *gorm.Config {
	return &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Duration(300) * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
			ParameterizedQueries:      true,
		},
	)}
}
