package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"golang-project-layout/config"
	"golang-project-layout/static"
)

type connection struct {
	dsn       string
	appConfig *config.AppConfig
	instance  *gorm.DB
}

func NewDBConn(dsn string, appConfig *config.AppConfig) config.DBConn {
	return &connection{dsn: dsn, appConfig: appConfig}
}

func (c *connection) DataSourceName() string {
	return c.dsn
}

func (c *connection) Open() (*gorm.DB, error) {
	if c.appConfig == nil {
		return nil, static.ErrNoAppConfig
	}

	var err error
	c.instance, err = gorm.Open(mysql.Open(c.dsn), c.appConfig.GormConfig)
	if nil != err {
		return nil, err
	}

	instanceDb, err := c.instance.DB()
	if nil != err {
		return nil, err
	}

	if c.appConfig.GormConfig.MaxOpenConnections > 0 {
		instanceDb.SetMaxOpenConns(c.appConfig.GormConfig.MaxOpenConnections)
	}

	if c.appConfig.GormConfig.MaxIdleConnections > 0 {
		instanceDb.SetMaxIdleConns(c.appConfig.GormConfig.MaxIdleConnections)
	}

	instanceDb.SetConnMaxLifetime(c.appConfig.GormConfig.ConnectionMaxTime)

	if c.appConfig.GormConfig.ConnectionIdleTime > 0 {
		instanceDb.SetConnMaxIdleTime(c.appConfig.GormConfig.ConnectionIdleTime)
	}

	return c.instance, nil
}

func (c *connection) Close() error {
	if c.instance == nil {
		return static.ErrUninitializedDatabase
	}

	gormDb, err := c.instance.DB()
	if err != nil {
		return err
	}

	return gormDb.Close()
}

func (c *connection) Instance() (*gorm.DB, error) {
	if c.instance == nil {
		return nil, static.ErrUninitializedDatabase
	}

	return c.instance, nil
}

func (c *connection) Ping() error {
	instance, err := c.Instance()
	if err != nil {
		return err
	}

	gormDb, err := instance.DB()
	if err != nil {
		return err
	}

	return gormDb.Ping()
}
