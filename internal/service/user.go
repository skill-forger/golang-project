package service

import (
	"context"

	"golang-project-layout/internal/contract"
	"golang-project-layout/internal/model"
	"golang-project-layout/internal/repository"

	"gorm.io/gorm"
)

type UserSvcImpl struct {
	db       *gorm.DB
	userRepo repository.UserRepo
}

func NewUserSvc(db *gorm.DB, userRepo repository.UserRepo) *UserSvcImpl {
	return &UserSvcImpl{
		db:       db,
		userRepo: userRepo,
	}
}

func (svc *UserSvcImpl) CreateUser(ctx context.Context, userDTO contract.CreatedUserDTO) (model.User, error) {
	panic("implement me")
}
