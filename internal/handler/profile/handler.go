package profile

import (
	"github.com/labstack/echo/v4"

	"golang-project-layout/server"

	hdl "golang-project-layout/internal/handler"
	svc "golang-project-layout/internal/service"
)

type handler struct {
	route      string
	profileSvc svc.Profile
}

func NewHandler(route string, profileSvc svc.Profile) hdl.Profile {
	return &handler{
		route:      route,
		profileSvc: profileSvc,
	}
}

func (h handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: h.route,
		Register: func(group *echo.Group) {
			group.GET("/", h.Get)
		},
	}
}

func (h handler) Get(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
