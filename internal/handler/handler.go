package handler

import (
	"github.com/labstack/echo/v4"

	"golang-project-layout/server"
)

type ResourceHandler interface {
	RegisterRoutes() server.HandlerRegistry
}

type Authentication interface {
	ResourceHandler
	SignIn(echo.Context) error
}

type Profile interface {
	ResourceHandler
	Get(echo.Context) error
}
