package Profile

import (
	repo "golang-project-layout/internal/repository"
	svc "golang-project-layout/internal/service"
)

type service struct {
	userRepo repo.User
}

func NewService(userRepo repo.User) svc.Profile {
	return &service{
		userRepo: userRepo,
	}
}

func (p *service) GetDetail() {
	//TODO implement me
	panic("implement me")
}
