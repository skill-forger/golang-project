package authentication

import (
	"github.com/labstack/echo/v4"

	"golang-project-layout/server"

	hdl "golang-project-layout/internal/handler"
	svc "golang-project-layout/internal/service"
)

type handler struct {
	route   string
	authSvc svc.Authentication
}

func NewHandler(route string, authSvc svc.Authentication) hdl.Authentication {
	return &handler{
		route:   route,
		authSvc: authSvc,
	}
}

func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: h.route,
		Register: func(group *echo.Group) {
			group.POST("/sign-in", h.SignIn)
		},
	}
}

func (h *handler) SignIn(e echo.Context) error {
	//TODO implement me
	panic("implement me")
}
