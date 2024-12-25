package users

import (
	"context"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(ctx context.Context, db *gorm.DB, user User) error
}

type UserRepoImpl struct{}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

func (r *UserRepoImpl) Create(ctx context.Context, db *gorm.DB, user User) error {
	result := db.WithContext(ctx).Create(&user)

	return result.Error
}
