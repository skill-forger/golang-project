package user

import (
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

func (r *repository) Read(id int) (*model.User, error) {
	var result *model.User

	query := r.db.Model(&model.User{}).First(&result, "`id` = ?", id)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Insert(o *model.User) (*model.User, error) {
	return nil, nil
}

func (r *repository) Update(o *model.User) (*model.User, error) {
	return nil, nil
}

func (r *repository) ReadByCondition(m map[string]any) (*model.User, error) {
	return nil, nil
}
