package repository

import (
	"context"

	"golang-project-layout/internal/model"
)

type User interface {
	Read(context.Context, int) (*model.User, error)
	Insert(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) (*model.User, error)
	ReadByCondition(map[string]any) (*model.User, error)
}
