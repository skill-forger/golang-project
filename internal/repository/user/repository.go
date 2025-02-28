package user

import (
	"context"

	"gorm.io/gorm"

	"golang-project-layout/internal/model"
	repo "golang-project-layout/internal/repository"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.User {
	return &repository{db: db}
}

func (r repository) Read(ctx context.Context, i int) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Insert(ctx context.Context, user *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) ReadByCondition(m map[string]any) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
