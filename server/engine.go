package server

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Engine interface {
	Address() string
	Startup(handlers ...HandlerRegistry) error
	Shutdown(ctx context.Context) error
}

type engine struct {
	address string
	server  *echo.Echo
}

func NewEngine(address string, configs ...ConfigProvider) Engine {
	echoServer := echo.New()

	for _, provide := range configs {
		provide(echoServer)
	}

	return &engine{address: address, server: echoServer}
}

func (e *engine) Address() string {
	return e.address
}

func (e *engine) Startup(handlers ...HandlerRegistry) error {
	if e.Address() == "" {
		return ErrMissingServerAddress
	}

	for _, handler := range handlers {
		handler.Register(e.server.Group(handler.Route))
	}

	return e.server.Start(e.Address())
}

func (e *engine) Shutdown(ctx context.Context) error {
	if e.server == nil {
		return ErrUninitializedEngine
	}

	return e.server.Shutdown(ctx)
}
