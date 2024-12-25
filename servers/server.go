package servers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"golang-project-layout/infra"
	"golang-project-layout/v1/users"
	"net/http"
)

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
func InitServer(app *infra.AppConfig) {
	e := echo.New()
	db, err := app.DB.Instance()
	if err != nil {
		panic(err)
	}

	// TODO: auto fill bearer token with EchoWrapHandler and swagger setup javascript file
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepo := users.NewUserRepo()
	userSvc := users.NewUserSvc(db, userRepo)
	userCtrl := users.NewUserCtrl(userSvc)

	apiGroup := e.Group("/api")
	v1Group := apiGroup.Group("/v1")

	// healthcheck
	v1Group.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// non authenticated, authorized routes
	// users
	usersGroup := v1Group.Group("/users")
	{
		usersGroup.POST("", userCtrl.CreateUser)
	}

	// jwt
	e.Use()

	// authenticated, authorized routes
	// users APIs
	users.RegisterUsersRoutes(usersGroup, userCtrl)

	e.Logger.Fatal(e.Start(app.ServerPort))
}
