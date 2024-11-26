package infra

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

type AppConfig struct {
	ServerPort string
	Env        string
	GormConfig *GormConfig
	DB         DBConn
}

func NewAppConfig() *AppConfig {
	appConfig := AppConfig{}

	// config server
	appConfig.ServerPort = viper.GetString("SERVER_PORT")

	// config gorm
	gormConfig := newDefaultConfig()
	maxOpenConnections := viper.GetInt("MAX_OPEN_CONNECTIONS")
	maxIdleConnections := viper.GetInt("MAX_IDLE_CONNECTIONS")
	gormConfig.MaxOpenConnections = maxOpenConnections
	gormConfig.MaxIdleConnections = maxIdleConnections
	gormConfig.ConnectionMaxTime = time.Hour * 99999
	gormConfig.ConnectionIdleTime = time.Hour * 99999
	appConfig.GormConfig = gormConfig

	// config database
	dbHost := viper.GetString("DB_HOST")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbPort := viper.GetInt("DB_PORT")
	dbConn := NewDBConn(fmt.Sprintf("%[1]s:%[2]s@tcp(%[3]s:%[4]d)/%[5]s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,     // 1
		dbPassword, // 2
		dbHost,     // 3
		dbPort,     // 4
		dbName,     // 5
	), &appConfig)

	appConfig.DB = dbConn

	_, err := dbConn.Open()
	if err != nil {
		log.Fatalln(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &appConfig
}
