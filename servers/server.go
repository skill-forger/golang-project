package servers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/swaggo/echo-swagger"
	"golang-project-layout/config"
	"golang-project-layout/database"
	"golang-project-layout/internal/controllers"
	"golang-project-layout/internal/repositories"
	"golang-project-layout/internal/services"
	"log"
	"net/http"
	"time"
)

type DataBaseConfig interface {
}

func NewAppConfig() *config.AppConfig {
	appConfig := config.AppConfig{}

	// config server
	appConfig.ServerPort = viper.GetString("SERVER_PORT")

	// config gorm
	gormConfig := database.NewDefaultConfig()
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
	dbConn := database.NewDBConn(fmt.Sprintf("%[1]s:%[2]s@tcp(%[3]s:%[4]d)/%[5]s?charset=utf8mb4&parseTime=True&loc=Local",
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

type UserSvc interface {
	CreateUser(ctx context.Context) error
}

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func InitServer(app *config.AppConfig) {
	e := echo.New()
	db, err := app.DB.Instance()
	if err != nil {
		panic(err)
	}

	// TODO: auto fill bearer token with EchoWrapHandler and swagger setup javascript file
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepo := repositories.NewUserRepo()
	userSvc := services.NewUserSvc(db, userRepo)
	userCtrl := controllers.NewUserCtrl(userSvc)

	apiGroup := e.Group("/api")

	// healthcheck
	apiGroup.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// non authenticated, authorized routes
	// users
	usersGroup := apiGroup.Group("/users")
	{
		usersGroup.POST("", userCtrl.CreateUser)
		usersGroup.GET("/:id", userCtrl.GetUser)
	}

	// authentication, authorization middlewares
	//e.Use()

	e.Logger.Fatal(e.Start(app.ServerPort))
}
