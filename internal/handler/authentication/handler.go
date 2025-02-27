package authentication

import (
	"github.com/labstack/echo/v4"

	hdl "golang-project-layout/internal/handler"
	svc "golang-project-layout/internal/service"
)

type handler struct {
	authSvc svc.Authentication
}

func NewHandler(authSvc svc.Authentication) hdl.Authentication {
	return &handler{
		authSvc: authSvc,
	}
}

func (h handler) RegisterRoutes() {
	//TODO implement me
	panic("implement me")
}

func (h handler) SignIn(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
