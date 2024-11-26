package users

import (
	"context"
	_ "golang-project-layout/infra"
	"gorm.io/gorm"
)

type UserSvcImpl struct {
	db       *gorm.DB
	userRepo UserRepo
}

func NewUserSvc(db *gorm.DB, userRepo UserRepo) *UserSvcImpl {
	return &UserSvcImpl{
		db:       db,
		userRepo: userRepo,
	}
}

func (svc *UserSvcImpl) CreateUser(ctx context.Context, userDTO CreatedUserDTO) (User, error) {
	panic("implement me")
}
