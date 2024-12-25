package users

import (
	"github.com/labstack/echo/v4"
)

func RegisterUsersRoutes(root *echo.Group, userCtrl UserCtrl) {
	root.GET("/:id", userCtrl.GetUser)
}
