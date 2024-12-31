package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang-project-layout/internal/dtos"
	"golang-project-layout/internal/models"
	"net/http"
)

type UserSvc interface {
	CreateUser(ctx context.Context, userDTO dtos.CreatedUserDTO) (models.User, error)
}

type UserCtrl interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
}

type UserCtrlImpl struct {
	UserSvc UserSvc
}

func NewUserCtrl(userSvc UserSvc) UserCtrl {
	return &UserCtrlImpl{
		UserSvc: userSvc,
	}
}

// @Summary     Create user
// @Description Create user
// @Accept      json
// @Produce     json
// @Param       json body     users.CreatedUserDTO  true  "User information for creating user"
// @Success     201  {object} users.CreatedUserDTO
// @Failure     400  "error"
// @Router      /api/v1/users [post]
func (ctrl *UserCtrlImpl) CreateUser(c echo.Context) error {
	panic("implement me")
}

// @Security    BearerAuth
// @param       Authorization header string true "Authorization"
// @Summary     Create user
// @Description Create user
// @Accept      json
// @Produce     json
// @Param       id path int true "1"
// @Success     200  {object} users.UserDetailDTO
// @Failure     400  "error"
// @Router      /api/v1/users/{id} [get]
func (ctrl *UserCtrlImpl) GetUser(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, dtos.UserDetailDTO{
		Username: "username" + id,
		Email:    "email@gmail.com",
		Password: "emailQwe123!@#",
	})
}
