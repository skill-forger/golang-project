package config

import (
	"gorm.io/gorm"
	"time"
)

type AppConfig struct {
	ServerPort string
	Env        string
	GormConfig *GormConfig
	DB         DBConn
}

type GormConfig struct {
	*gorm.Config
	MaxOpenConnections int
	MaxIdleConnections int
	ConnectionMaxTime  time.Duration
	ConnectionIdleTime time.Duration
}

type DBConn interface {
	DataSourceName() string
	Open() (*gorm.DB, error)
	Close() error
	Instance() (*gorm.DB, error)
	Ping() error
}
