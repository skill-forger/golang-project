package services

import (
	"context"
	"golang-project-layout/internal/dtos"
	"golang-project-layout/internal/models"
	"golang-project-layout/internal/repositories"
	"gorm.io/gorm"
)

type UserSvcImpl struct {
	db       *gorm.DB
	userRepo repositories.UserRepo
}

func NewUserSvc(db *gorm.DB, userRepo repositories.UserRepo) *UserSvcImpl {
	return &UserSvcImpl{
		db:       db,
		userRepo: userRepo,
	}
}

func (svc *UserSvcImpl) CreateUser(ctx context.Context, userDTO dtos.CreatedUserDTO) (models.User, error) {
	panic("implement me")
}
