package health

import (
	"gorm.io/gorm"

	"golang-project-layout/internal/handler"
	hdl "golang-project-layout/internal/handler/health"
)

func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, db)
}
