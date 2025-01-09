package repositories

import (
	"context"
	"golang-project-layout/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(ctx context.Context, db *gorm.DB, user models.User) error
}

type UserRepoImpl struct{}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

func (r *UserRepoImpl) Create(ctx context.Context, db *gorm.DB, user models.User) error {
	result := db.WithContext(ctx).Create(&user)

	return result.Error
}
