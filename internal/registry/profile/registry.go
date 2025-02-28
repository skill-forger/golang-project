package profile

import (
	"gorm.io/gorm"

	"golang-project-layout/internal/handler"
	hdl "golang-project-layout/internal/handler/profile"
	repo "golang-project-layout/internal/repository/user"
	svc "golang-project-layout/internal/service/profile"
)

func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(repo.NewRepository(db)))
}
