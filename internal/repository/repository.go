package repository

import (
	"golang-project-layout/internal/model"
)

type User interface {
	Read(int) (*model.User, error)
	Insert(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	ReadByCondition(map[string]any) (*model.User, error)
}
