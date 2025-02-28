package handler

import (
	"github.com/labstack/echo/v4"

	"golang-project-layout/server"
)

// ResourceHandler             represents all API resource handler
// @title                      golang project layout server swagger API
// @version                    1.0
// @description                This is the swagger API for golang project layout.
// @termsOfService             http://swagger.io/terms/
// @contact.name               API Support
// @contact.url                http://www.swagger.io/support
// @contact.email              support@swagger.io
// @license.name               Apache 2.0
// @license.url                http://www.apache.org/licenses/LICENSE-2.0.html
// @host                       localhost:3000
// @BasePath                   /
// @securityDefinitions.basic  BearerToken
// @externalDocs.description   OpenAPI
// @externalDocs.url           https://swagger.io/resources/open-api/
type ResourceHandler interface {
	RegisterRoutes() server.HandlerRegistry
}

// Authentication represents all authentication resource handler
type Authentication interface {
	ResourceHandler
	SignIn(echo.Context) error
}

// Profile represents all profile resource handler
type Profile interface {
	ResourceHandler
	Get(echo.Context) error
}
