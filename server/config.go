package server

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var (
	ErrMissingServerAddress = errors.New("server address is missing")
	ErrUninitializedEngine  = errors.New("server engine is not initialized")
)

type ConfigProvider func(*echo.Echo)

type HandlerRegistry struct {
	Route    string
	Register func(*echo.Group)
}
