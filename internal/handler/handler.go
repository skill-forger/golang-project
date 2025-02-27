package handler

import (
	"github.com/labstack/echo/v4"
)

type BaseHandler interface {
	RegisterRoutes()
}

type Authentication interface {
	BaseHandler
	SignIn(echo.Context) error
}

type Profile interface {
	BaseHandler
	Get(echo.Context) error
}
