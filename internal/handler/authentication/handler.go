package authentication

import (
	"net/http"

	"github.com/labstack/echo/v4"

	ct "golang-project-layout/internal/contract"
	hdl "golang-project-layout/internal/handler"
	svc "golang-project-layout/internal/service"
	"golang-project-layout/server"
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
	request := new(ct.SignInRequest)
	if err := e.Bind(request); err != nil {
		return err
	}

	if err := e.Validate(request); err != nil {
		return err
	}

	response, err := h.authSvc.SignIn(request)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}
