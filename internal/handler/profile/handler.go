package profile

import (
	"net/http"

	"github.com/labstack/echo/v4"

	hdl "golang-project-layout/internal/handler"
	svc "golang-project-layout/internal/service"
	"golang-project-layout/server"
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

func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route:           h.route,
		IsAuthenticated: true,
		Register: func(group *echo.Group) {
			group.GET("", h.Get)
		},
	}
}

// Get   handles the profile detail request
//	@Summary		Respond profile detail information
//	@Description	Respond profile detail information
//	@Tags			profile
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Success		200	{object}	contract.ProfileResponse
//	@Failure		400	{object}	error
//	@Router			/profile [get]
func (h *handler) Get(e echo.Context) error {
	ctxUser, err := hdl.GetContextUser(e)
	if err != nil {
		return err
	}

	response, err := h.profileSvc.GetByID(ctxUser.ID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}
