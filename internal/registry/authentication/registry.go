package authentication

import (
	"gorm.io/gorm"

	"golang-project-layout/internal/handler"
	hdl "golang-project-layout/internal/handler/authentication"
	repo "golang-project-layout/internal/repository/user"
	svc "golang-project-layout/internal/service/authentication"
	"golang-project-layout/util/hashing"
)

// NewRegistry returns new resource handler for authentication API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(repo.NewRepository(db), hashing.NewBcrypt()))
}
