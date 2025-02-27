package profile

import (
	hdl "golang-project-layout/internal/handler"
	svc "golang-project-layout/internal/service"

	"github.com/labstack/echo/v4"
)

type handler struct {
	userSvc svc.User
}

func NewHandler(userSvc svc.User) hdl.Profile {
	return &handler{
		userSvc: userSvc,
	}
}

func (h handler) RegisterRoutes() {
	//TODO implement me
	panic("implement me")
}

func (h handler) Get(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
